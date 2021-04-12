package ddos

import (
	"errors"
	"log"
	"net"
	"net/http"
	"sync"
	"time"

	"golang.org/x/time/rate"
)

// Struct that define a visitor
// limiter contains the number of request he can do
// lastSeen is the timestamp of his last request
type visitor struct {
	limiter  *rate.Limiter
	lastSeen time.Time
}

// Struct that define the DDOSProtection
// globalLimiter contains the number of request available for the RVP
// visitors is the map of visitor
// mu is the mutex for writing in the visitor map
// tokenRefreshRatePerSecond refresh request number rate per second
// tokenBucketSize number of request available for a new visitor
type DDOSProtection struct {
	typeProtection            string
	globalLimiter             *rate.Limiter
	visitors                  map[string]*visitor
	mu                        sync.Mutex
	tokenRefreshRatePerSecond rate.Limit
	tokenBucketSize           int
}

// We create the DDOSProtection struct and initialize it with the parameters.
// If the protection type is UserLimit we start the cleanupVisitors thread
func New(typeProtection string, tokenRefreshRatePerSecond rate.Limit, tokenBucketSize int, refreshVisitors int) *DDOSProtection {
	ddosProtection := new(DDOSProtection)
	ddosProtection.typeProtection = typeProtection
	ddosProtection.globalLimiter = rate.NewLimiter(tokenRefreshRatePerSecond, tokenBucketSize)
	ddosProtection.tokenRefreshRatePerSecond = tokenRefreshRatePerSecond
	ddosProtection.tokenBucketSize = tokenBucketSize

	if typeProtection == "UserLimit" {
		go ddosProtection.cleanupVisitors(refreshVisitors)
		ddosProtection.visitors = make(map[string]*visitor)
	}

	return ddosProtection
}

// Function that check if the user is already save in the map and save his new timestamp
func (dd *DDOSProtection) getVisitor(ip string) *rate.Limiter {
	dd.mu.Lock()
	defer dd.mu.Unlock()

	v, exists := dd.visitors[ip]
	if !exists {
		limiter := rate.NewLimiter(dd.tokenRefreshRatePerSecond, dd.tokenBucketSize)
		// Include the current time when creating a new visitor.
		dd.visitors[ip] = &visitor{limiter, time.Now()}
		return limiter
	}

	// Update the last seen time for the visitor.
	v.lastSeen = time.Now()
	return v.limiter
}

// Every minute check the map for visitors that haven't been seen for
// more than 3 minutes and delete the entries.
func (dd *DDOSProtection) cleanupVisitors(refreshTime int) {
	for {
		time.Sleep(time.Minute)

		dd.mu.Lock()
		for ip, v := range dd.visitors {
			if time.Since(v.lastSeen) > time.Duration(refreshTime)*time.Second {
				delete(dd.visitors, ip)
			}
		}
		dd.mu.Unlock()
	}
}

// Check if the user is allow to send a new request with his limiter
func (dd *DDOSProtection) userLimit(r *http.Request) error {
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		log.Println(err.Error())
		return errors.New("InternalServerError")
	}

	limiter := dd.getVisitor(ip)
	if !limiter.Allow() {
		return errors.New("StatusTooManyRequests")
	}

	return nil
}

// Check if the user is allow to send a new request with the global limiter
func (dd *DDOSProtection) globalLimit() error {
	if !dd.globalLimiter.Allow() {
		return errors.New("StatusTooManyRequests")
	}

	return nil
}

// this function call the verification method based on the type protection
func (dd *DDOSProtection) CheckLimit(r *http.Request) error {
	if dd.typeProtection == "UserLimit" {
		return dd.userLimit(r)
	} else {
		return dd.globalLimit()
	}
}

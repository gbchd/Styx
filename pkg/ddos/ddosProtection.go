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

type visitor struct {
	limiter  *rate.Limiter
	lastSeen time.Time
}

type DDOSProtection struct {
	typeProtection            string
	globalLimiter             *rate.Limiter
	visitors                  map[string]*visitor
	mu                        sync.Mutex
	tokenRefreshRatePerSecond rate.Limit
	tokenBucketSize           int
}

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

func (dd *DDOSProtection) globalLimit() error {
	if !dd.globalLimiter.Allow() {
		return errors.New("StatusTooManyRequests")
	}

	return nil
}

func (dd *DDOSProtection) CheckLimit(r *http.Request) error {
	if dd.typeProtection == "UserLimit" {
		return dd.userLimit(r)
	} else {
		return dd.globalLimit()
	}
}

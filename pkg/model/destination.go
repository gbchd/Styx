package model

import (
	"net/url"
	"sync"
)

// Struct that contains all the settings for one destination
// URL is the url of the destination
// Alive is the status of the destination
// Weight Importance pour le load balancing
// Mux Mutex to manage writting concurence
type Destination struct {
	URL    *url.URL
	Alive  bool
	Weight int64
	Mux    *sync.RWMutex
}

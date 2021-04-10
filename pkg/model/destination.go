package model

import (
	"net/url"
	"sync"
)

type Destination struct {
	URL    *url.URL
	Alive  bool
	Weight int64
	Mux    *sync.RWMutex
}

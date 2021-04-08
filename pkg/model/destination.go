package model

import (
	"net/url"
	"sync"
)

type Destination struct {
	URL    *url.URL
	Alive  bool
	weight int
	mux    sync.RWMutex
}

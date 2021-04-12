package rvp

import (
	"net/http"
	"net/http/httputil"

	"github.com/guillaumebchd/styx/pkg/model"
)

// GenerateProxy is a function that will take a conf object and return a reverse proxy correctly configured.
func GenerateProxy(conf *model.ReverseProxy) httputil.ReverseProxy {
	rvp := httputil.ReverseProxy{
		Director: func(r *http.Request) {

			// We get the destination
			site := conf.Sites[r.Host]
			destination := site.Get()

			r.Header.Add("X-Forwarded-Host", r.Host)
			r.Header.Add("X-Origin-Host", destination.URL.Host)
			r.Host = destination.URL.Host
			r.URL.Host = destination.URL.Host
			r.URL.Scheme = destination.URL.Scheme

			// We could handle the override of the path here
			// r.URL.RawPath = destination.URL.RawPath

		},
	}

	return rvp

}

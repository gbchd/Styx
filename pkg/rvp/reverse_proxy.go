package rvp

import (
	"net/http"
	"net/http/httputil"

	"github.com/guillaumebchd/styx/pkg/conf"
)

func getDestination(host string, conf conf.Sites) string {
	// We get the correct destination from the Configuration struct
	routes, ok := conf.Map[host]
	if ok {
		// TODO :
		// We implement the load balancer function here to determine which backend we should use.
		// For now we will take the first in the array
		return routes.Destinations[0]
	}

	// If we don't find the host in our configuration we redirect to the default destination
	return conf.Default
}

// GenerateProxy is a function that will take a conf object and return a reverse proxy correctly configured.
func GenerateProxy(conf conf.Sites) httputil.ReverseProxy {
	rvp := httputil.ReverseProxy{
		Director: func(r *http.Request) {

			// We get the destination
			destination := getDestination(r.Host, conf)

			r.Header.Add("X-Forwarded-Host", r.Host)
			r.Header.Add("X-Origin-Host", destination)
			r.Host = destination
			r.URL.Host = destination
			r.URL.Scheme = "https"

			// We could handle the override of the path here

		},
	}

	return rvp

}

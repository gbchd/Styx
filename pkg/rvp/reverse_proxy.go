package rvp

import (
	"net/http"
	"net/http/httputil"
)

// Configuration is a simple struct that contains the routes (in a map) and the default route
// in case we don't find our host in the map.
// FIXME : This struct is temporary and will be replace by the real struct in t he configuration package.
type Configuration struct {
	Default string
	Routes  map[string]Route
}

// Route is a struct that holds an entrypoint and multiple destinations.
// Its role is to store a link for our reverse proxy. An entrypoint is the url the client type that will send them to the reverse proxy,
// the destinations are the backend server that the reverse proxy will redirect the request to.
type Route struct {
	Entrypoint   string
	Destinations []string

	// We could add some attributes to overrides the header etc...
	// but let's keep it simple for now
}

func getDestination(host string, conf Configuration) string {
	// We get the correct destination from the Configuration struct
	routes, ok := conf.Routes[host]
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
func GenerateProxy(conf Configuration) httputil.ReverseProxy {
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

// GenerateTestConfiguration is a test function that create a Configuration struct.
func GenerateTestConfiguration() Configuration {

	route1 := Route{
		Entrypoint:   "localhost:80",
		Destinations: []string{"google.com"},
	}
	route2 := Route{
		Entrypoint:   "127.0.0.1:80",
		Destinations: []string{"twitter.com"},
	}

	routes := map[string]Route{
		route1.Entrypoint: route1,
		route2.Entrypoint: route2,
	}

	conf := Configuration{
		Default: "google.com",
		Routes:  routes,
	}

	return conf
}

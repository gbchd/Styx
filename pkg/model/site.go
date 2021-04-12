package model

// Struct for our reverse proxy
// Default is our default destination
// Sites is a map of all our destination
type ReverseProxy struct {
	Default Destination
	Sites   map[string]*Site
}

// Site is the struct of one destination for our RVP
// Name is the name of the destination
// Entrypoint is the url of the entrypoint
// DestinationsPool is all the destination to redirect
type Site struct {
	Name       string
	Entrypoint string
	DestinationsPool
}

// Function that create our ReverseProxy struct
func Create(def Destination, sites map[string]*Site) *ReverseProxy {
	return &ReverseProxy{
		Default: def,
		Sites:   sites,
	}
}

func Get() *Destination {
	return nil
}

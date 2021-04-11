package model

type ReverseProxy struct {
	Default Destination
	Sites   map[string]*Site
}

type Site struct {
	Name       string
	Entrypoint string
	DestinationsPool
}

func Create(def Destination, sites map[string]*Site) *ReverseProxy {
	return &ReverseProxy{
		Default: def,
		Sites:   sites,
	}
}

func Get() *Destination {
	return nil
}

package model

type ReverseProxy struct {
	Default string
	Sites   map[string]Site
}

type Site struct {
	Entrypoint string
	DestinationsPool
}

func Create() *ReverseProxy {
	return nil
}

func Get() *Destination {
	return nil
}

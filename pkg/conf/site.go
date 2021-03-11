package conf

import (
	"github.com/pelletier/go-toml"
)

// Site is a struct that represent a possible destination for the RVP
// It contains the name of the site we want to access
// The entrypoint url that the user is going to try to access
// The addresses of the possible destinations
type Site struct {
	Name         string
	Entrypoint   string
	Destinations []string
}

// Sites is a struct that contains the default route in case we can't access the normal destination
// and a list of site that the RVP can access
type Sites struct {
	Default string
	Map     map[string]Site // map[Entrypoint]Site

	// We could add some attributes to overrides the header etc...
	// but let's keep it simple for now
}

// GetSites gets all the sites that the rvp can reach from the configuration file
func GetSites(config *toml.Tree) Sites {

	siteTree := config.Get("sites").([]*toml.Tree)
	def := config.Get("Default").(*toml.Tree)

	sitesMap := make(map[string]Site)

	for index := range siteTree {

		name := siteTree[index].Get("name").(string)
		entrypoint := siteTree[index].Get("entrypoint").(string)
		destinations := siteTree[index].GetArray("addresses").([]string)

		site := Site{
			Name:         name,
			Entrypoint:   entrypoint,
			Destinations: destinations,
		}
		sitesMap[siteTree[index].Get("entrypoint").(string)] = site
	}

	sites := Sites{
		Default: def.Get("default_route").(string),
		Map:     sitesMap,
	}

	return sites
}

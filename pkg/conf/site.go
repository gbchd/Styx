package conf

import (
	"github.com/guillaumebchd/styx/pkg/entities"
	"github.com/pelletier/go-toml"
)

// Site is a struct that represent a possible destination for the RVP
// It contains the name of the site we want to access
// The entrypoint url that the user is going to try to access
// The addresses of the possible destinations
type Site struct {
	Name       string
	Route      entities.Route
	Entrypoint string        // deprecated
	Addresses  []interface{} // deprecated
}

// Sites is a struct that contains the default route in case we can't access the normal destination
// and a list of site that the RVP can access
type Sites struct {
	Default  string
	SiteList map[string]Site
}

// GetSites gets all the sites that the rvp can reach from the configuration file
func GetSites(config *toml.Tree) Sites {

	siteTree := config.Get("sites").([]*toml.Tree)
	def := config.Get("Default").(*toml.Tree)

	list := make(map[string]Site)

	for index := range siteTree {
		site := Site{
			Name: siteTree[index].Get("name").(string),
			Route: entities.Route{
				Entrypoint:   siteTree[index].Get("entrypoint").(string),
				Destinations: siteTree[index].Get("addresses").([]string),
			},
			Entrypoint: siteTree[index].Get("entrypoint").(string),
			Addresses:  siteTree[index].Get("addresses").([]interface{}),
		}
		list[siteTree[index].Get("name").(string)] = site
	}

	sites := Sites{
		Default:  def.Get("default_route").(string),
		SiteList: list,
	}

	return sites
}

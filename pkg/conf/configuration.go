package conf

import (
	"github.com/pelletier/go-toml"
)

// ServerConfiguration is a simple struct that contains the parameters of the RVP server
// it contains the url and the port of the server
type ServerConfiguration struct {
	Port int64
	ServerName string
}

// Site is a struct that represent a possible destination for the RVP
// It contains the name of the site we want to access
// The entrypoint url that the user is going to try to access
// The addresses of the possible destinations
type Site struct {
	Name string
	Entrypoint string
	Addresses []interface{}
}

// Sites is a struct that contains the default route in case we can't access the normal destination
// and a list of site that the RVP can access
type Sites struct {
	Default string
	SiteList map[string]Site
}

// DDosConfiguration contains all the parameters that we can need for our DDOS protection
type DDosConfiguration struct {
	MaxRequest int64
	MaxRequestPerUser int64
	VerificationTimer int64
}

// This method gets the rvp server informations from the configuration file
func GetServerConfiguration() ServerConfiguration {

	config, _ := toml.LoadFile("../../pkg/conf/configuration.toml")

	conf := ServerConfiguration{
		Port: config.Get("server.port").(int64),
		ServerName: config.Get("server.server_name").(string),
	}

	return conf
}	

// This method gets all the sites that the rvp can reach from the configuration file
func GetSites() Sites {

	config, _ := toml.LoadFile("../../pkg/conf/configuration.toml")

	siteTree := config.Get("sites").([]*toml.Tree)
	def := config.Get("Default").(*toml.Tree)

	list := make(map[string]Site)

	for index := range siteTree {
		site := Site{
			Name: siteTree[index].Get("name").(string),
			Entrypoint: siteTree[index].Get("entrypoint").(string),
			Addresses: siteTree[index].Get("addresses").([]interface{}),
		}
		list[siteTree[index].Get("name").(string)] = site
	}

	sites := Sites{
		Default: def.Get("default_route").(string),
		SiteList: list,
	}

	return sites
}

// This method gets all the parameters that are going to be useful for the DDos protection
// of the rvp
func GetDDosConfiguration() DDosConfiguration {

	config, _ := toml.LoadFile("../../pkg/conf/configuration.toml")

	ddosConfig := config.Get("DDOS_Parameters").(*toml.Tree)

	conf := DDosConfiguration{
		MaxRequest: ddosConfig.Get("max_request").(int64),
		MaxRequestPerUser: ddosConfig.Get("max_request_per_user").(int64),
		VerificationTimer: ddosConfig.Get("verification_timer").(int64),
	}

	return conf
}
package conf

import (
	"github.com/pelletier/go-toml"

	"github.com/guillaumebchd/styx/pkg/model"
)

// Configuration holds all the parameters defined in our configuration file.
type Configuration struct {
	Server       Server
	ReverseProxy *model.ReverseProxy
	DDos         DDos
}

// Get returns a Configuration object from a toml.tree
func Get(config *toml.Tree) Configuration {
	server := GetServer(config)
	def, sites := GetSites(config)
	ddos := GetDDos(config)

	reverseProxy := model.Create(def, sites)

	return Configuration{
		Server:       server,
		ReverseProxy: reverseProxy,
		DDos:         ddos,
	}
}

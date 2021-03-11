package conf

import "github.com/pelletier/go-toml"

// Configuration holds all the parameters defined in our configuration file.
type Configuration struct {
	Server Server
	Sites  Sites
	DDos   DDos
}

// Get returns a Configuration object from a toml.tree
func Get(config *toml.Tree) Configuration {
	server := GetServer(config)
	sites := GetSites(config)
	ddos := GetDDos(config)

	return Configuration{
		Server: server,
		Sites:  sites,
		DDos:   ddos,
	}
}

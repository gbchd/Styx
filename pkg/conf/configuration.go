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

// Host is a struct that represent a possible destination for the RVP
// It contains the name of the host we want to access
// The entrypoint url that the user is going to try to access
// The IP of the destination
// The port of the destination
type Host struct {
	Name string
	Entrypoint string
	Ip string
	Port int64
}

// Hosts is a struct that contains the default route in case we can't access the normal destination
// and a list of host that the RVP can access
type Hosts struct {
	Default string
	HostList map[string]Host
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

// This method gets all the hosts that the rvp can reach from the configuration file
func GetHosts() Hosts {

	config, _ := toml.LoadFile("../../pkg/conf/configuration.toml")

	hostTree := config.Get("hosts").([]*toml.Tree)
	def := config.Get("Default").(*toml.Tree)

	list := make(map[string]Host)

	for index := range hostTree {
		host := Host{
			Name: hostTree[index].Get("name").(string),
			Entrypoint: hostTree[index].Get("entrypoint").(string),
			Ip: hostTree[index].Get("ip").(string),
			Port: hostTree[index].Get("port").(int64),
		}
		list[hostTree[index].Get("name").(string)] = host
	}

	hosts := Hosts{
		Default: def.Get("default_route").(string),
		HostList: list,
	}

	return hosts
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
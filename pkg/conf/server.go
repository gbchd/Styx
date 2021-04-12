package conf

import "github.com/pelletier/go-toml"

// Server is a simple struct that contains the parameters of the RVP server
// it contains the url and the port of the server
type Server struct {
	Port       int64
	ServerName string
}

// GetServer gets the rvp server informations from the configuration file
func GetServer(config *toml.Tree) Server {

	conf := Server{
		Port:       config.Get("server.port").(int64),
		ServerName: config.Get("server.server_name").(string),
	}

	return conf
}

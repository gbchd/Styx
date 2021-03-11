package conf

import "github.com/pelletier/go-toml"

// DDos contains all the parameters that we can need for our DDOS protection
type DDos struct {
	MaxRequest        int64
	MaxRequestPerUser int64
	VerificationTimer int64
}

// GetDDos gets all the parameters that are going to be useful for the DDos protection
// of the rvp
func GetDDos(config *toml.Tree) DDos {

	ddosConfig := config.Get("DDOS_Parameters").(*toml.Tree)

	conf := DDos{
		MaxRequest:        ddosConfig.Get("max_request").(int64),
		MaxRequestPerUser: ddosConfig.Get("max_request_per_user").(int64),
		VerificationTimer: ddosConfig.Get("verification_timer").(int64),
	}

	return conf
}

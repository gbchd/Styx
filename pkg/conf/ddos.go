package conf

import "github.com/pelletier/go-toml"

// DDos contains all the parameters that we can need for our DDOS protection
type DDos struct {
	RefreshRequestRate int
	MaxRequestPerUser  int
	VerificationTimer  int
}

// GetDDos gets all the parameters that are going to be useful for the DDos protection
// of the rvp
func GetDDos(config *toml.Tree) DDos {

	ddosConfig := config.Get("DDOS_Parameters").(*toml.Tree)

	conf := DDos{
		MaxRequestPerUser:  ddosConfig.Get("max_request_per_user").(int),
		RefreshRequestRate: ddosConfig.Get("refresh_request_rate").(int),
		VerificationTimer:  ddosConfig.Get("verification_timer").(int),
	}

	return conf
}

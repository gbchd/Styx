package main

import (
	"fmt"
	
	"github.com/guillaumebchd/styx/pkg/conf"
)

func main() {
	server := conf.GetServerConfiguration()
	fmt.Println(server.Port)
	fmt.Println(server.ServerName)

	ddos := conf.GetDDosConfiguration()
	fmt.Println(ddos.MaxRequest)
	fmt.Println(ddos.MaxRequestPerUser)
	fmt.Println(ddos.VerificationTimer)

	sites := conf.GetSites()
	fmt.Println(sites)
	fmt.Println(sites.SiteList["google"].Addresses[0])
}

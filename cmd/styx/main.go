package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/guillaumebchd/styx/pkg/conf"
	"github.com/guillaumebchd/styx/pkg/rvp"
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

	// We create our configuration
	conf := rvp.GenerateTestConfiguration()

	r := mux.NewRouter()

	// We create our reverse proxy from our configuration object object
	reverseProxy := rvp.GenerateProxy(conf)

	// We capture all the paths and we redirect it to the reverse proxy
	r.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reverseProxy.ServeHTTP(w, r)
	})

	srv := &http.Server{
		Handler:      r,
		Addr:         "0.0.0.0:80",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Println("Starting server on : " + srv.Addr)
	log.Fatal(srv.ListenAndServe())

}

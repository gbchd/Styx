package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/guillaumebchd/styx/pkg/conf"
	"github.com/guillaumebchd/styx/pkg/ddos"
	"github.com/guillaumebchd/styx/pkg/rvp"
	"github.com/pelletier/go-toml"
	"golang.org/x/time/rate"
)

func main() {

	// We read our configuration file
	configuration, err := toml.LoadFile("./configuration.toml")
	if err != nil {
		log.Fatal(err)
	}
	conf := conf.Get(configuration)

	// We create our router
	r := mux.NewRouter()

	// We create our reverse proxy from our configuration object object
	reverseProxy := rvp.GenerateProxy(conf.ReverseProxy)

	// We capture all the paths and we redirect it to the reverse proxy
	r.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reverseProxy.ServeHTTP(w, r)
	})

	if conf.DDos.Activate {
		ddosProtect := ddos.New(conf.DDos.Type, rate.Limit(conf.DDos.RefreshRequestRate), conf.DDos.MaxRequestPerUser, conf.DDos.VerificationTimer)
		r.Use(ddosProtect.Proctection)
	}

	// We set this value to 10000 to be able to test with wrk.
	// Normally, we should keep the default value of 2 but we can't do our tests with it.
	http.DefaultTransport.(*http.Transport).MaxIdleConns = 10000
	http.DefaultTransport.(*http.Transport).MaxIdleConnsPerHost = 10000

	srv := &http.Server{
		Handler:      r,
		Addr:         fmt.Sprintf("0.0.0.0:%d", conf.Server.Port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Println("Starting server " + conf.Server.ServerName + " on : " + srv.Addr)
	log.Fatal(srv.ListenAndServe())
}

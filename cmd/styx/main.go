package main

import (
	"log"
	"net/http"

	"golang.org/x/time/rate"
)

var token_bucket_size int = 3
var token_refresh_rate_per_second rate.Limit = 1

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", okHandler)

	// Wrap the servemux with the limit middleware.
	log.Println("Listening on :4000...")
	//http.ListenAndServe(":4000", global_limit(mux))
	http.ListenAndServe(":4000", user_limit(mux))
}

func okHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}
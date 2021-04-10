package main

import (
    "net/http"

    "golang.org/x/time/rate"
)

var limiter = rate.NewLimiter(token_refresh_rate_per_second, token_bucket_size)

func global_limit(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        if limiter.Allow() == false {
            http.Error(w, http.StatusText(429), http.StatusTooManyRequests)
            return
        }

        next.ServeHTTP(w, r)
    })
}
package ddos

import (
	"log"
	"net/http"
)

func (ddosProtect *DDOSProtection) Proctection(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		result := ddosProtect.CheckLimit(r)

		if result == nil {
			next.ServeHTTP(w, r)
		} else {
			log.Printf(result.Error())
			switch result.Error() {
			case "InternalServerError":
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			case "StatusTooManyRequests":
				http.Error(w, http.StatusText(http.StatusTooManyRequests), http.StatusTooManyRequests)

			}
		}
	})
}

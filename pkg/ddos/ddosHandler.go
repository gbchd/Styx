package ddos

import (
	"net/http"
)

// Handler that check if the user is enable to send his request with the ddos protection
// If there is an Internal Error or the Status Too Many Request we notify the user
func (ddosProtect *DDOSProtection) Proctection(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		result := ddosProtect.CheckLimit(r)

		if result == nil {
			next.ServeHTTP(w, r)
		} else {
			switch result.Error() {
			case "InternalServerError":
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			case "StatusTooManyRequests":
				http.Error(w, http.StatusText(http.StatusTooManyRequests), http.StatusTooManyRequests)

			}
		}
	})
}

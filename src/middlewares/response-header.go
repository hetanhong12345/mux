package middlewares

import (
	"net/http"
)

func ResponseHeader(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", origin)
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		// post method call ParseForm
		if r.Method != "GET" {
			r.ParseForm()
		}

		next.ServeHTTP(w, r)
	})
}

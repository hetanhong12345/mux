package middlewares

import (
	"encoding/json"
	"log"
	"net/http"
)

var unauthorized = ErrorMessage{"401", "unauthorized"}

func LoginRequried(next http.Handler) http.Handler {
	result, _ := json.Marshal(unauthorized)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		cookie, err := r.Cookie("mobile")
		if err != nil {
			log.Printf("<----login-required middleware request url is %s ----->\n", r.RequestURI)
			w.Write(result)
			return
		}

		log.Printf("<----login-required middleware cookie  is %s ----->\n", cookie)
		next.ServeHTTP(w, r)
	})
}

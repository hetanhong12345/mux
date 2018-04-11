package middlewares

import (
	"net/http"
	"fmt"
	"encoding/json"
)


var unauthorized = ErrorMessage{"401", "unauthorized"}

func LoginRequried(next http.Handler) http.Handler {
	result, _ := json.Marshal(unauthorized)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		cookie, err := r.Cookie("auth")
		if err != nil {
			fmt.Printf("<----login-required middleware request url is %s ----->\n", r.RequestURI)
			w.Write(result)
		}

		fmt.Printf("<----login-required middleware cookie  is %s ----->\n", cookie)
		next.ServeHTTP(w, r)
	})
}

package middlewares

import (
	"net/http"
	"fmt"
	"encoding/json"
)

var recovery = ErrorMessage{"500", "Internal Server Error"}

func Recovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rcr := recover(); rcr != nil {
				recovery.msg = fmt.Sprintf("%v", rcr)
				result, _ := json.Marshal(recovery)
				w.Write(result)
			}

		}()

		next.ServeHTTP(w, r)
	})
}

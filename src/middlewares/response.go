package middlewares

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var recovery = ErrorMessage{"500", "Internal Server Error"}

func Response(w http.ResponseWriter, result interface{}) {
	if rcy := recover(); rcy != nil {
		recovery.Msg = fmt.Sprintf("%v", rcy)
		response, _ := json.Marshal(recovery)
		w.Write(response)
		return
	}
	response, _ := json.Marshal(result)
	w.Write(response)

}

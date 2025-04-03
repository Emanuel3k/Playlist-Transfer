package response

import (
	"encoding/json"
	"net/http"
)

func Send(w http.ResponseWriter, code int, body any) {
	if body == nil {
		w.WriteHeader(code)
		return
	}

	bytes, err := json.Marshal(body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(code)

	_, err = w.Write(bytes)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

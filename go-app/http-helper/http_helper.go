package http_helper

import (
	"encoding/json"
	"net/http"
)

type M map[string]interface{}

func WriteJson(w http.ResponseWriter, val interface{}) {
	data, err := json.Marshal(val)
	if err != nil {
		WriteError(w, err, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func WriteError(w http.ResponseWriter, err error, errorStatusCode int) {
	http.Error(w, err.Error(), errorStatusCode)
}

package http_helper

import (
	"encoding/json"
	"net/http"
)

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

func GetQueryParams(r *http.Request) (map[string]string, error) {
	var values map[string][]string
	m := make(map[string]string)
	switch r.Method {
	case http.MethodGet:
		values = r.URL.Query()
	case http.MethodPut, http.MethodPost, http.MethodPatch:
		err := r.ParseForm()
		if err != nil {
			return nil, err
		}
		values = r.PostForm
	}
	for k, v := range values {
		m[k] = v[0]
	}
	return m, nil
}

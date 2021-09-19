package helper

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func WriteJson(w http.ResponseWriter, val interface{}) {
	data, err := json.Marshal(val)
	if err != nil {
		WriteError(w, err, http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(data)
}

func WriteError(w http.ResponseWriter, err error, errorStatusCode int) {
	http.Error(w, err.Error(), errorStatusCode)
}

func ProcessJsonBody(r *http.Request) (map[string]string, error) {
	m := make(map[string]interface{})
	decoder := json.NewDecoder(r.Body)
	defer dclose(r.Body)
	err := decoder.Decode(&m)
	if err != nil {
		return nil, err
	}
	result, err := toMapStringString(m)
	return result, err
}

func toMapStringString(m map[string]interface{}) (map[string]string, error) {
	res := make(map[string]string)

	for k, v := range m {
		if val, ok := v.(map[string]interface{}); ok {
			bytes, err := json.Marshal(val)
			if err != nil {
				return nil, err
			}
			res[k] = string(bytes)
		} else {
			res[k] = fmt.Sprintf("%v", v)
		}
	}
	return res, nil
}

func dclose(c io.Closer) {
	if err := c.Close(); err != nil {
		log.Println(err)
	}
}

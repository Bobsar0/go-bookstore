package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Unmarshal json data from request body into [data]
func ParseBody(r *http.Request, data interface{}) {
	if body, err := ioutil.ReadAll(r.Body); err == nil {

		// parse the JSON-encoded body and store the result in data
		if err := json.Unmarshal([]byte(body), data); err != nil {
			return
		}
	}
}

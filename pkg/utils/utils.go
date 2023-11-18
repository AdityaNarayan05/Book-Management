package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// ParseBody reads the request body from an HTTP request and unmarshals it into the provided interface.
func ParseBody(r *http.Request, x interface{}) {
	// Read the request body
	if body, err := ioutil.ReadAll(r.Body); err == nil {
		// Unmarshal the JSON data into the provided interface
		if err := json.Unmarshal(body, x); err != nil {
			// If there is an error in unmarshaling, simply return
			return
		}
	}
}

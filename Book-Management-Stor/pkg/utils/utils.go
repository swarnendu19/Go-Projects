package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

// Importing necessary packages for handling JSON and HTTP requests

// ParseBody func parses the request body into the provided interface.
// Reads the request body and attempts to unmarshal it into the provided interface.
// If successful, assigns the parsed data to the provided interface.

func ParseBody(r *http.Request, x interface{}) {
	body, err := io.ReadAll(r.Body)
	if err == nil {
		err := json.Unmarshal([]byte(body), x)
		if err != nil {
			return
		}
	}
}

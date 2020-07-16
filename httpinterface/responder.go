package httpinterface

import (
	"encoding/json"
	"net/http"
)

// Just alias for the data map
type (
	data map[string]interface{}
)

// Predefined headers
const (
	ContentTypeHeader = "Content-Type"
	ContentTypeJSON   = "application/json"
)

// Serve data as JSON as response
func jsonResponse(w http.ResponseWriter, status int, v interface{}) error {
	w.Header().Set(ContentTypeHeader, ContentTypeJSON)
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

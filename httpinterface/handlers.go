package httpinterface

import (
	"encoding/json"
	"net/http"
)

type (
	// Function to validate json data and parse into given interface
	binder func(req *http.Request, data interface{}, rul, msg map[string][]string) error

	interactor interface {
		Exec(a, b, c bool, d float64, e, f int) (string, float64, error)
	}

	// Options structure
	Options struct {
		A bool    `json:"a"`
		B bool    `json:"b"`
		C bool    `json:"c"`
		D float64 `json:"d"`
		E int     `json:"e"`
		F int     `json:"f"`
	}
)

// Interactor handler
func Interactor(i interactor) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := &Options{}
		err := json.NewDecoder(r.Body).Decode(data)
		if err != nil {
			jsonResponse(w, http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
			return
		}

		h, k, err := i.Exec(data.A, data.B, data.C, data.D, data.E, data.F)
		if err != nil {
			jsonResponse(w, http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
			return
		}

		jsonResponse(w, http.StatusOK, map[string]interface{}{"h": h, "k": k})
	}
}

// Root path handler
func Root(buildTag string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		jsonResponse(w, http.StatusOK, map[string]interface{}{"build_tag": buildTag})
	}
}

// HealthCheck handler
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
}

// NotFound handler
func NotFound(w http.ResponseWriter, r *http.Request) {
	jsonResponse(w, http.StatusNotFound, map[string]interface{}{"error": http.StatusText(http.StatusNotFound)})
}

// MethodNotAllowed handler
func MethodNotAllowed(w http.ResponseWriter, r *http.Request) {
	jsonResponse(w, http.StatusMethodNotAllowed, map[string]interface{}{"error": http.StatusText(http.StatusMethodNotAllowed)})
}

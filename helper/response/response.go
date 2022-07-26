package helper

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Status bool        `json:"status"`
	Data   interface{} `json:"data,omitempty"`
	Error  string      `json:"error,omitempty"`
}

func BuildResponse(w http.ResponseWriter, status bool, code int, data interface{}, errors string) {
	resp := &Response{
		Status: status,
		Data:   data,
		Error:  errors,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(resp)
}

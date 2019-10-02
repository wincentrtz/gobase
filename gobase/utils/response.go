package utils

import (
	"encoding/json"
	"net/http"
)

func WriteResponse(w http.ResponseWriter, value interface{}) {
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(value)
}

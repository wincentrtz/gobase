package utils

import (
	"encoding/json"
	"net/http"

	"github.com/wincentrtz/gobase/models/responses"
)

func SuccessResponse(w http.ResponseWriter, value interface{}) {
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(value)
}

func ErrorResponse(w http.ResponseWriter, err responses.ErrorResponse) {
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(err)
}

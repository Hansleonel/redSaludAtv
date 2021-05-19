package utils

import (
	"encoding/json"
	"net/http"
)

func RespondWithSuccess(data interface{}, w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

func RespondWithError(err error, w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(err.Error())
}

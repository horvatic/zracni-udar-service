package controller

import (
	"encoding/json"
	"net/http"

	"github.com/horvatic/zracni-udar-service/pkg/service"
)

func sendCreateResult(w http.ResponseWriter, errorType service.ErrorType, err error) {
	if errorType == service.JsonError || errorType == service.BadRequest {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	} else if errorType == service.DatabaseError {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func sendJson(w http.ResponseWriter, j interface{}, errorType service.ErrorType, err error) {
	if errorType == service.JsonError || errorType == service.BadRequest {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	} else if errorType == service.DatabaseError {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(j)
}

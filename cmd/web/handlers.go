package main

import (
	"encoding/json"
	"net/http"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	type response struct {
		Message string `json:"message"`
	}

	message := response{Message: "Welcome home"}
	// jsonMessage, err := json.Marshal(message)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// }
	json.NewEncoder(w).Encode(message)
}

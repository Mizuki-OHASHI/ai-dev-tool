package controller

import (
	"encoding/json"
	"net/http"
	"main/usecase"
)

func HandleRequest(w http.ResponseWriter, r *http.Request) {
	// Parse request data
	var requestData map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Validate request data
	if requestData["key"] == nil {
		http.Error(w, "key is missing", http.StatusBadRequest)
		return
	}

	// Call the appropriate usecase function with the parsed data
	responseData, err := usecase.SomeFunction(requestData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Handle the response from the usecase and send it back to the client
	jsonResponse, err := json.Marshal(responseData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}
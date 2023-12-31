package controller

import (
	"encoding/json"
	"main/model"
	"main/usecase"
	"net/http"

	"github.com/gorilla/mux"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	logRequest(r)
	users, err := usecase.GetUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(users)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	logRequest(r)
	params := mux.Vars(r)
	id := params["userId"]
	if id == "" {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	user, err := usecase.GetUser(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	logRequest(r)
	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = usecase.CreateUser(&user)
	if err != nil {
		if err.Error() == "Username already exists" {
			http.Error(w, err.Error(), http.StatusConflict)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	logRequest(r)
	params := mux.Vars(r)
	id := params["userId"]
	if id == "" {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user.ID = id
	err = usecase.UpdateUser(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	logRequest(r)
	params := mux.Vars(r)
	id := params["userId"]
	if id == "" {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	err := usecase.DeleteUser(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

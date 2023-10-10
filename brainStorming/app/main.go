package main

import (
	"fmt"
	"log"
	"main/dao"
	"main/controller"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	dao.CloseDBWithSysCall()

	r := mux.NewRouter()
	r.HandleFunc("/users", controller.GetUsers).Methods("GET")
	r.HandleFunc("/users", controller.CreateUser).Methods("POST")
	r.HandleFunc("/users/{userId}", controller.GetUser).Methods("GET")
	r.HandleFunc("/users/{userId}", controller.UpdateUser).Methods("PUT")
	r.HandleFunc("/users/{userId}", controller.DeleteUser).Methods("DELETE")

	log.Printf("Listening...")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), r); err != nil {
		log.Fatal(err)
	}
}

func init() {

	if err := godotenv.Load(".env"); err != nil {
		log.Printf("fail: godotenv.Load, %v\n", err)
	}

	dao.OpenSql()
}

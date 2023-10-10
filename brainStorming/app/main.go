package main

import (
	"fmt"
	"log"
	"main/controller"
	"main/dao"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	dao.CloseDBWithSysCall()

	r := mux.NewRouter()
	r.HandleFunc("/user", controller.GetUser).Methods("GET")

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

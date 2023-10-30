package controller

import (
	"log"
	"net/http"
)

func logRequest(r *http.Request) {
	log.Printf("Received request: %s %s", r.Method, r.URL)
}

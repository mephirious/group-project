package main

import (
	"log"
	"net/http"

	"github.com/mephirious/group-project/services/customer/internal/handlers"
)

func main() {
	r := http.NewServeMux()

	r.HandleFunc("POST /customers", handlers.RegisterCustomer)
	r.HandleFunc("POST /customers/login", handlers.LoginCustomer)
	// r.HandleFunc("GET /customers/me", handlers.)

	log.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", r)
}

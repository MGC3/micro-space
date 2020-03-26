package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	// create router
	r := mux.NewRouter()

	// grab port from env
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("Error setting port")
	}

	// start server
	fmt.Println("Server listening at:", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}

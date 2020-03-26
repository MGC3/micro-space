package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("Error setting port")
	}

	fmt.Println("Server listening at:", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

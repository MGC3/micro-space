package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = "8080"

func main() {
	fmt.Println("Server listening at:", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

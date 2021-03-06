package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

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

	// create custom server, for more control
	srv := http.Server{
		Addr:         ":" + port,
		Handler:      r,
		ReadTimeout:  1 * time.Second,   // max duration for reading client req - make higher if client uploading big file, etc.
		WriteTimeout: 1 * time.Second,   // max duration for writing to client - small is fine for sending basic JSON
		IdleTimeout:  120 * time.Second, // all around connection pooling - allows client to resuse the same connection, useful if they have multiple requests. Particularly useful for microservices
	}

	// start server - use goroutine so it doesn't block
	go func() {
		log.Println("Server listening at:", port)
		log.Fatal(srv.ListenAndServe())
	}()

	// listen for SIGINT (ctrl + c) command
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	// block channel until SIGINT received
	<-c

	// attempt to wait 30sec for current connections to finish before shutting down server
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	srv.Shutdown(ctx)
	log.Println("gracefully shutdown")
}

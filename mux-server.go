package main

import (
	"net/http"
	"router"
	"time"
	"log"
)

func main() {
	r := router.Root

	srv := &http.Server{
		Handler: r,
		Addr:    "0.0.0.0:3002",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}


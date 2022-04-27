package main

import (
	"log"
	"net/http"
	"time"

	"lightsaid.com/millionare/cmd/api/handlers"
	"lightsaid.com/millionare/cmd/api/routes"
	"lightsaid.com/millionare/internal/repository"
)

var httpAddr = ":4000"

func main() {
	repo := repository.NewRepository()

	handler := handlers.NewAPIHandler(repo)

	r := routes.NewRoutes(handler)

	s := &http.Server{
		Addr:           httpAddr,
		Handler:        r,
		ReadTimeout:    20 * time.Second,
		WriteTimeout:   20 * time.Second,
		MaxHeaderBytes: 2 * 1024,
	}

	log.Println("Starting HTTP Server on ", httpAddr)

	err = s.ListenAndServe()

	if err != nil {
		log.Println("HTTP Server Start Failed: ", err)
	}
}

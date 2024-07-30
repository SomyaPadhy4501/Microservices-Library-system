package main

import (
	"log"
	"net/http"
	"time"

	"github.com/SomyaPadhy4501/book-store/pkg/routes"
	"github.com/gorilla/mux"
)

func RouterMethod() {
	r := mux.NewRouter()
	s := r.PathPrefix("/book").Subrouter()
	routes.StoreRoutes(s)
	r.Use(loggingMiddleware)
	s.Use(AuthMiddleware)
	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}

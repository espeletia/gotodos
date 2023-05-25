package main

import (
	"net/http"
	"todos/handlers"

	"github.com/XSAM/otelsql"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	dbconn, err := otelsql.Open("postgres", "postgres://postgres:postgres@localhost:5434/postgres?sslmode=disable")
	if err != nil {
		panic(err)
	}

	router := mux.NewRouter()
	router.HandleFunc("/users", handlers.CreateUser(dbconn)).Methods("POST")
	router.HandleFunc("/users/{id}", handlers.GetUser(dbconn)).Methods("GET")
	api := http.Server{
		Addr:         "0.0.0.0:" + "8888",
		Handler:      router,
	}
	api.ListenAndServe()
}

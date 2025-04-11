package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github.com/AspenFresh/lab4-webapp/internal"
	"github.com/AspenFresh/lab4-webapp/internal/adapters/postgres"
	"github.com/AspenFresh/lab4-webapp/internal/ports/rest"
)

func main() {
	db, err := sqlx.Connect("postgres", "user=postgres password=nik10sen dbname=lab4db sslmode=disable")
	if err != nil {
		log.Fatal("failed to connect to DB:", err)
	}
	defer db.Close()

	dbClient := postgres.NewClient(db)
	service := internal.NewUserService(dbClient)
	handler := rest.NewHandler(service)

	r := mux.NewRouter()
	r.HandleFunc("/users", handler.CreateUserHandler).Methods("POST")

	log.Println("Server is running on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

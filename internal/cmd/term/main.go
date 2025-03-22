package main

import (
	"context"
	"github.com/DenisGoldiner/webapp/internal/ports/ftp"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"

	"github.com/DenisGoldiner/webapp/internal"
	"github.com/DenisGoldiner/webapp/internal/adapters/postgres"
)

func main() {
	run()
}

func run() {
	dbExec, err := newDB()
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	travellersClient := postgres.NewClient(dbExec)
	travellersService := internal.NewTravellers(travellersClient)
	travellersParser := ftp.NewParser(travellersService)

	if err = travellersParser.Run(ctx, "/Users/denys/Go/src/github.com/DenisGoldiner/webapp/internal/integration/data/test_1.csv"); err != nil {
		log.Printf("error running travellers import: %v", err)
		return
	}
}

func newDB() (sqlx.ExtContext, error) {
	dsn := "postgres://postgres:postgres@localhost:5432/travellers?sslmode=disable"
	conn, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

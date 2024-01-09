package main

import (
	"context"
	"hello/db"
	"log"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

func run() error {
	ctx := context.Background()

	d, err := pgxpool.New(ctx, "user=dbuser password=rachit dbname=testdb sslmode=disable host=localhost port=5432")
	if err != nil {
		return err
	}

	q := db.New(d)

	insertedAuthor, err := q.CreateAuthor(ctx, db.CreateAuthorParams{
		Name: "Mr Beans",
		Bio: pgtype.Text{
			String: "The funniest man alive",
			Valid:  true,
		},
	})
	if err != nil {
		return err
	}

	log.Println(insertedAuthor)
	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

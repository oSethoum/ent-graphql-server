package db

import (
	"context"
	"entql/ent"

	"entgo.io/ent/dialect/sql/schema"
	_ "github.com/mattn/go-sqlite3"
)

func Init() *ent.Client {
	db, err := ent.Open("sqlite3", "file:db.sqlite?_fk=1")

	if err != nil {
		panic(err)
	}

	err = db.Schema.Create(context.Background(), schema.WithDropColumn(true), schema.WithDropIndex(true))

	if err != nil {
		panic(err)
	}

	return db
}

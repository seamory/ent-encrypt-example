package db

import (
	"context"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"main/ent"
)

var Client *ent.Client

func init() {
	Client = connect()
}

func connect() *ent.Client {
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	//defer client.Close()
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	return client
}

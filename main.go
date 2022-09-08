package main

import (
	"context"
	"fmt"
	"log"
	"main/internal/cipher"
	"main/internal/db"
	"main/internal/property"
)

func main() {
	cipher.LoadKeySets("./ks.bin")
	// add user
	ctx := context.Background()
	for i := 0; i < 10; i++ {
		username := fmt.Sprintf("test0%d", i)
		password := fmt.Sprintf("test0%dPassword", i)
		db.Client.AccountInfo.Create().SetUsername(username).SetPassword(property.Password(password)).Save(ctx)
	}
	// print the raw value from table account_infos
	queryContext, err := db.Client.QueryContext(ctx, "select * from account_infos", nil)
	if err != nil {
		log.Fatalln(err)
	}
	type row struct {
		id       int
		username string
		password string
	}
	for queryContext.Next() {
		r := &row{}
		err := queryContext.Scan(&r.id, &r.username, &r.password)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(r)
	}
	// print the value through ent from table account_infos
	table, err := db.Client.AccountInfo.Query().All(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	for _, info := range table {
		fmt.Println(info)
	}
}

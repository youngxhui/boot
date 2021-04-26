package db

import (
	"boot/ent"
	"context"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var client *ent.Client

func init() {
	var err error
	client, err = ent.Open("mysql", "root:1234@tcp(localhost:3307)/boot?parseTime=True")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}

	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

}

package db

import (
	"boot/ent"
	"context"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
	"log"
)

// DB 数据库连接对象
var DB *gorm.DB

var client *ent.Client

func init() {
	//connStr := "postgres://kong:kong@localhost:5432/boot?sslmode=disable"
	// mysql://localhost:3306/boot?serverTimezone=UTC
	//client, err := ent.Open("postgres","host=localhost port=5432 user=kong dbname=boot password=kong sslmode=disable")
	var err error
	client, err = ent.Open("mysql", "root:1234@tcp(localhost:3306)/boot?parseTime=True")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}

	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

}

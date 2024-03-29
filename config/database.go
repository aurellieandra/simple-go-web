package config

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDB() {
	db, err := sql.Open("mysql", "root:admin123@/go_sql")
	if err != nil {
		panic(err)
	}

	log.Println("Database is connected...")
	DB = db
}

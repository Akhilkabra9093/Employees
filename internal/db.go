package internal

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var Db *sql.DB

func InitDB() {
	dataSourceName := "root:Akhil@9093@tcp(localhost:3306)/employees"
	var err error
	Db, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	log.Println("Connected to database")
}

func CloseDB() {
	if Db != nil {
		Db.Close()
		log.Println("Database connection closed")
	}
}

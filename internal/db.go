package internal

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"sync"
	"time"
)

var (
	db   *sql.DB
	once sync.Once
)

func InitDB() {
	dataSourceName := "root:Akhil@9093@tcp(localhost:3306)/employees"
	var err error
	db, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)
	log.Println("Connected to database")
}

func CloseDB() {
	if db != nil {
		db.Close()
		log.Println("Database connection closed")
	}
}

func GetDB() *sql.DB {
	return db
}

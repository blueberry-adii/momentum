package database

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Connect() {
	username := "root"
	pass := "pass"
	host := "localhost"
	port := "3306"
	dbname := "momentum"

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", username, pass, host, port, dbname)
	var err error

	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Error connecting to DB: ", err)
	}

	DB.SetMaxOpenConns(25)
	DB.SetMaxIdleConns(25)
	DB.SetConnMaxLifetime(5 * time.Minute)

	if err := DB.Ping(); err != nil {
		log.Fatal("Database not reachable: ", err)
	}

	log.Println("MySQL connected successfully!!!")
}

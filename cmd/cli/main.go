package main

import (
	"log"
	"os"
)

import (
	"database/sql"
	_ "github.com/lib/pq"
)

func main() {
	url := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", url)
	if err != nil {
		log.Fatalln("db conn", url, "error", err)
	}

	log.Println("db open", db)

	err = db.Ping()
	if err != nil {
		log.Fatalln("db ping error", err)
	}
}

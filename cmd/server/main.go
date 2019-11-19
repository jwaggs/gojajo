package main

import (
	"fmt"
	"github.com/jwaggs/piggy"
	"log"
	"net/http"
	"os"
)

import (
	"database/sql"
	_ "github.com/lib/pq"
)

func main() {
	log.Println("entering server")
	defer log.Println("exiting server")

	// create our db connection pool
	_ = openDB()
	// listen for and serve http requests
	serve()
}

// getEnv is a simple wrapper to return a fallback value if os env var doesn't exist
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

// openDB configures the postgres database connection pool and returns a sql.DB reference
func openDB() *sql.DB {
	url := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", url)
	if err != nil {
		log.Fatalln("db conn:", url, "error", err)
	}
	// since sql.Open only creates the conn pool, force successful contact with the db server by pinging it.
	if err = db.Ping(); err != nil {
		log.Fatalln("db ping error:", err)
	}
	log.Println("db open:", db)
	return db
}

// serve configures and runs the server
func serve() {
	port := getEnv("PORT", "3000")
	addr := fmt.Sprintf(":%s", port)
	log.Println("serving on", addr)

	r := piggy.Router()
	err := http.ListenAndServe(addr, r)
	if err != nil {
		log.Fatalln("http serve error:", err)
	}
}
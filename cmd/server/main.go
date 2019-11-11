package main

import (
	"fmt"
	"github.com/go-chi/chi"
	"log"
	"net/http"
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

	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("piggy"))
	})

	port := os.Getenv("PORT")
	if port == "" {
		// default if no env var exists
		port = "3000"
	}

	addr := fmt.Sprintf(":%s", port)
	log.Println("serving on", addr)
	http.ListenAndServe(addr, r)
}

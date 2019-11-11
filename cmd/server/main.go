package main

import (
	"fmt"
	"github.com/go-chi/chi"
	"net/http"
	"os"
)

func main() {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("piggy"))
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	addr := fmt.Sprintf(":%s", port)
	http.ListenAndServe(addr, r)
}

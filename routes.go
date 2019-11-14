package piggy

import (
	"fmt"
	"github.com/go-chi/chi"
	"log"
	"net/http"
)

// Router wraps the http router creation and returns a reference to the new *chi.Mux router
func Router() *chi.Mux {
	// the router we will attach all "routes" to
	r := chi.NewRouter()

	// root route just writes piggy
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		rsp, err := w.Write([]byte("piggy"))
		log.Println("piggy rsp:", rsp, "err:", err)
	})

	var counter int
	// route to add 1 to our in-memory counter. (GET request because browser demo)
	r.Get("/add", func(w http.ResponseWriter, r *http.Request) {
		counter += 1
		rsp, err := w.Write([]byte(fmt.Sprintf("counter = %d", counter)))
		log.Println("add rsp:", rsp, "err:", err)
	})
	// route to subtract 1 from our in-memory counter. (GET request because browser demo)
	r.Get("/sub", func(w http.ResponseWriter, r *http.Request) {
		counter -= 1
		rsp, err := w.Write([]byte(fmt.Sprintf("counter = %d", counter)))
		log.Println("sub rsp:", rsp, "err:", err)
	})

	return r
}
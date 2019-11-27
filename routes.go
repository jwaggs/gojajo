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

	var started, finished int
	// root route just writes piggy
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("piggy started:", started, "finished:", finished)
		rsp, err := w.Write([]byte(fmt.Sprintln("piggy started:", started, "finished:", finished)))
		log.Println("piggy rsp:", rsp, "err:", err)
	})
	// route to add 1 to our in-memory counter. (GET request because browser demo)
	r.Get("/started", func(w http.ResponseWriter, r *http.Request) {
		started += 1
		log.Println("scans started:", started)
		rsp, err := w.Write([]byte(fmt.Sprintf("started = %d", started)))
		log.Println("started rsp:", rsp, "err:", err)
	})
	// route to subtract 1 from our in-memory counter. (GET request because browser demo)
	r.Get("/finished", func(w http.ResponseWriter, r *http.Request) {
		finished += 1
		log.Println("scans finished:", finished)
		rsp, err := w.Write([]byte(fmt.Sprintf("finished = %d", finished)))
		log.Println("finished rsp:", rsp, "err:", err)
	})

	return r
}
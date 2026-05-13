package main

import (
	"fmt"
	"log"
	"net/http"
)

// * routes to be registered
func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, "<h1>Welcome to my awesome site!</h1>")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, `<h1>Contact Page</h1>
	<p>To get in touch, email me at <a href="mailto:moaz@innolabs.ai">moaz@innolabs.ai</a></p>`)
}

func readPathHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, r.URL.Path)
}

func main() {
	//* Define a new server
	mux := http.NewServeMux()

	//* register routes
	mux.HandleFunc("/", readPathHandler)
	// mux.HandleFunc("/contact", contactHandler)

	fmt.Println("starting server at :3000")

	//* starting the server
	err := http.ListenAndServe(":3000", mux)
	if err != nil {
		log.Fatalf("couldn't start the server.")
	}
}

package main

import (
	"fmt"
	"net/http"
)

// * a single route to be registered
func handlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Welcome to my awesome site!</h1>")
}

func main() {
	//* Define a new server
	mux := http.NewServeMux()

	//* register routes
	mux.HandleFunc("/", handlerFunc)

	fmt.Println("starting server at :3000")

	//* starting the server
	http.ListenAndServe(":3000", mux)
}

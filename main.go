package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	//* New chi Router
	r := chi.NewRouter()

	//* Setup Routes for chi
	setupRoutes(r)

	fmt.Println("starting server at :3000")

	//* starting the server
	err := http.ListenAndServe(":3000", r)
	if err != nil {
		log.Fatalf("couldn't start the server.")
	}
}

func setupRoutes(chi *chi.Mux) {
	chi.Get("/", homeHandler)
	chi.Get("/contact", contactHandler)
	chi.Get("/faq", faqHandler)
}

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

func faqHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, `
<h1>FAQ Page</h1>
<h3>Q: Is there a free version?</h3>
<p>A: Yes! We offer a free trial for 30 days on any paid plans.</p>

<h3>Q: What are your support hours?</h3>
<p>A: We have support staff answering emails 24/7, though response times may be a
bit slower on weekends.</p>

<h3>Q: How do I contact support?</h3>
<p>A: Email us - <a href="mailto:support@innolabs.ai">support@innolabs.ai</a></p>
	`)
}

package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	//* New chi Router
	r := chi.NewRouter()

	//* App Middleware "Logger"
	middleware.Logger(r)

	//* Setup Routes for chi
	setupRoutes(r)

	fmt.Println("starting server at :3000")

	//* starting the server
	err := http.ListenAndServe(":3000", r)
	if err != nil {
		log.Fatal("couldn't start the server:", err)
	}
}

func setupRoutes(chi *chi.Mux) {
	chi.Get("/", homeHandler)
	chi.Get("/contact", contactHandler)
	chi.Get("/faq", faqHandler)
	chi.Get("/metrics/{metricID}", metricsHandler)

}

func executeTemplate(w http.ResponseWriter, filePath string) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	t, err := template.ParseFiles(filePath)
	if err != nil {
		errMsg := "There was an error parsing the template"
		log.Printf("%ss: %v", errMsg, err)
		http.Error(w, errMsg, http.StatusInternalServerError)
		return
	}

	err = t.Execute(w, nil)
	if err != nil {
		errMsg := "There was an error executing the template"
		log.Printf("%ss: %v", errMsg, err)
		http.Error(w, errMsg, http.StatusInternalServerError)
		return
	}
}

// * routes to be registered
func homeHandler(w http.ResponseWriter, r *http.Request) {
	//* Go Template
	tPath := filepath.Join("templates", "home.gohtml")
	executeTemplate(w, tPath)

	//* Go Templ
	// templates.Home().Render(r.Context(), w)
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	//* Go Template
	tPath := filepath.Join("templates", "contact.gohtml")
	executeTemplate(w, tPath)

	//* Go Templ
	// templates.Contact().Render(r.Context(), w)
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

// Exercises
// using URL parameters
func metricsHandler(w http.ResponseWriter, r *http.Request) {
	metricID := chi.URLParam(r, "metricID")
	w.Header().Add("Content-Type", "text/html; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "<p>Requested Metric: %v</p>", metricID)
}

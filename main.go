package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/MeYo0o/lenslocked/templates"
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
	chi.Get("/exercise", exerciseHandler)
}

func executeTemplate(w http.ResponseWriter, filePath string, data any) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	t, err := template.ParseFiles(filePath)
	if err != nil {
		errMsg := "There was an error parsing the template"
		log.Printf("%ss: %v", errMsg, err)
		http.Error(w, errMsg, http.StatusInternalServerError)
		return
	}

	err = t.Execute(w, data)
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
	executeTemplate(w, tPath, nil)

	//* Go Templ
	// templates.Home().Render(r.Context(), w)
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	//* Go Template
	tPath := filepath.Join("templates", "contact.gohtml")
	executeTemplate(w, tPath, nil)

	//* Go Templ
	// templates.Contact().Render(r.Context(), w)
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	//* Go Template
	tPath := filepath.Join("templates", "faq.gohtml")
	executeTemplate(w, tPath, nil)

	//* Go Templ
	// templates.FAQ().Render(r.Context(), w)
}

func exerciseHandler(w http.ResponseWriter, r *http.Request) {
	dataList := struct {
		Items []string
	}{
		[]string{"M", "Y", "D", "A", "M"},
	}

	//* Go Template
	tPath := filepath.Join("templates", "exercise.gohtml")
	executeTemplate(w, tPath, dataList)

	//* Go Templ - i had to make sure the header Content-Type = "text/html" especially for this case because the html elements were being escaped.
	w.Header().Add("Content-Type", "text/html; charset=utf-8")
	templates.Exercise(dataList.Items).Render(r.Context(), w)
}

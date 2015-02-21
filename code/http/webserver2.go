package main

import (
	"bytes"
	"log"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
)

func main() {
	// We need to create a router
	router := mux.NewRouter().StrictSlash(true)
	// Add the "index" or root path
	router.HandleFunc("/", Index)
	// Fire up the server
	log.Println("Starting server on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

// This is the "index" handler
func Index(w http.ResponseWriter, r *http.Request) {
	// Fill out page data for index
	pageData := PageData{
		Title: "Index Page",
		Body:  "This is the body of the page.",
	}

	// Render a template with our page data
	tmpl, err := htmlTemplate(pageData)

	// If we got an error, write it out and exit
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	// All went well, so write out the template
	w.Write([]byte(tmpl))
}

// This is a basic struct to hold basic page data variables
type PageData struct {
	Title string
	Body  string
}

func htmlTemplate(pageData PageData) (string, error) {
	// Define a basic HTML template
	html := `<HTML>
	<head><title>{{.Title}}</title></head>
	<body>
	{{.Body}}
	</body>
	</HTML>`

	// Parse the template
	tmpl, err := template.New("index").Parse(html)

	// We need somewhere to write the executed template to
	var out bytes.Buffer
	// Render the template with the data we passed in
	err = tmpl.Execute(&out, pageData)
	// Return the template and the error
	return out.String(), err
}

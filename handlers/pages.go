package handlers

import (
	"html/template"
	"net/http"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "home.tmpl", nil)
}

func ScannerPage(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "scanner.tmpl", nil)
}

func RenderTemplate(w http.ResponseWriter, tmpl string, data map[string]interface{}) {
	t, err := template.ParseFiles("templates/" + tmpl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(w, data)
}

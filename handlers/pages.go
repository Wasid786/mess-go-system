package handlers

import (
	"fmt"
	"html/template"
	"net/http"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "home.tmpl")
}

// Scanner page handler
func ScannerPage(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "scanner.tmpl")
}

// RenderTemplate function
func RenderTemplate(w http.ResponseWriter, tmplName string) {
	tmpl, err := template.ParseFiles(
		"templates/home.tmpl",
		"templates/scanner.tmpl",
		"templates/checkstudents.tmpl",
	)
	if err != nil {
		http.Error(w, "Error loading template", http.StatusInternalServerError)
		fmt.Println("Template parsing error:", err)
		return
	}

	err = tmpl.ExecuteTemplate(w, tmplName, nil)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		fmt.Println("Template execution error:", err)
	}
}

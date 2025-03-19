package main

import (
	"fmt"
	"html/template"
	"messGo/database"
	"messGo/handlers"

	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Initialize database
	err := database.OpenDB("new_user:localhost@/mess?parseTime=true")
	if err != nil {
		fmt.Println("Failed to connect to database:", err)
		return
	}

	// server the page
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("templates/scan.tmpl")
		if err != nil {
			http.Error(w, "Error loading template", http.StatusInternalServerError)
			return

		}
		tmpl.Execute(w, nil)
	})

	// Define routes
	http.HandleFunc("/scan", handlers.ScanQRCode)

	// Start server
	fmt.Println("Server started at :8080")
	http.ListenAndServe(":8080", nil)
}

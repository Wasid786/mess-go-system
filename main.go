package main

import (
	"fmt"
	"messGo/database"
	"messGo/handlers"

	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	err := database.OpenDB("new_user:localhost@/mess?parseTime=true")
	if err != nil {
		fmt.Println("Failed to connect to database:", err)
		return
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handlers.RenderTemplate(w, "home.tmpl")
	})
	http.HandleFunc("/scanner", func(w http.ResponseWriter, r *http.Request) {
		handlers.RenderTemplate(w, "scanner.tmpl")
	})

	// Backend Route
	http.HandleFunc("/scan", handlers.ScanQRCode)

	// Start server
	fmt.Println("Server started at :8080")
	http.ListenAndServe(":8080", nil)
}

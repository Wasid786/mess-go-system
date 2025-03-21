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

	http.HandleFunc("/", handlers.HomePage)
	http.HandleFunc("/scanner", handlers.ScannerPage)
	http.HandleFunc("/check", handlers.GetStudents)

	// Post Route
	http.HandleFunc("/scan", handlers.ScanQRCode)
	http.HandleFunc("/students", handlers.CreateStudent)

	// Start server
	fmt.Println("Server started at :8080")
	http.ListenAndServe(":8080", nil)
}

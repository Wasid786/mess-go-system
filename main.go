package main

import (
	"fmt"
	"messGo/database"
	"messGo/handlers"

	"net/http"
)

func main() {
	// Initialize database
	err := database.InitDB("new_user:localhost@/mess?parseTime=true")
	if err != nil {
		fmt.Println("Failed to connect to database:", err)
		return
	}

	// Define routes
	http.HandleFunc("/scan", handlers.ScanQRCode)

	// Start server
	fmt.Println("Server started at :8080")
	http.ListenAndServe(":8080", nil)
}

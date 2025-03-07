package handlers

import (
	"database/sql"
	"encoding/json"
	"messGo/database"
	"messGo/models"
	"net/http"
	"time"
)

func ScanQRCode(w http.ResponseWriter, r *http.Request) {
	var req struct {
		StudentID string `json:"student_id"`
		MealType  string `json:"meal_type"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// Check if the student exists
	var student models.Student
	err := database.DB.QueryRow("SELECT id, name, hostel_id FROM students WHERE student_id = ?", req.StudentID).Scan(&student.ID, &student.Name, &student.HostelID)
	if err == sql.ErrNoRows {
		http.Error(w, "Student not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	// Check if the student has already taken the meal for the day
	var meal models.Meal
	today := time.Now().Format("2006-01-02")
	err = database.DB.QueryRow("SELECT id FROM meals WHERE student_id = ? AND meal_type = ? AND date = ?", req.StudentID, req.MealType, today).Scan(&meal.ID)
	if err == nil {
		http.Error(w, "Meal already taken", http.StatusForbidden)
		return
	} else if err != sql.ErrNoRows {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	// Allow the meal
	_, err = database.DB.Exec("INSERT INTO meals (student_id, meal_type, date) VALUES (?, ?, ?)", req.StudentID, req.MealType, today)
	if err != nil {
		http.Error(w, "Failed to record meal", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Meal allowed"})
}

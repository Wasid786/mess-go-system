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
	w.Header().Set("Content-Type", "application/json")

	var req struct {
		CourseID string `json:"course_id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid Request Format"})
		return
	}

	// Check if the student exists
	var student models.Student
	err := database.DB.QueryRow("SELECT name, hostel, course_id, enroll_no, registered_session, mess_slip FROM students WHERE course_id = ?", req.CourseID).Scan(&student.Name, &student.Hostel, &student.CourseID, &student.EnrollNo, &student.RegisteredSession, &student.MessSlip)
	if err == sql.ErrNoRows {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "Student Not Found"})
		return
	} else if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Dataabase error"})
		return
	}

	// Determine meal type based on time
	loc, _ := time.LoadLocation("Asia/Kolkata")
	currentTime := time.Now().In(loc)
	hour := currentTime.Hour()

	var mealType string
	switch {
	case (hour >= 6 && hour < 11):
		mealType = "Breakfast"
	case (hour >= 11 && hour < 18):
		mealType = "Lunch"
	case (hour >= 18 && hour < 23):
		mealType = "Dinner"
	default:
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(map[string]string{"error": "This is not meal time!"})
		return
	}

	// Check if the student has already taken the meal for the day
	var meal models.Meal
	today := currentTime.Format("2006-01-02")

	err = database.DB.QueryRow("SELECT id FROM meals WHERE course_id = ? AND meal_type = ? AND date = ?", req.CourseID, mealType, today).Scan(&meal.ID)
	if err == nil {
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(map[string]string{"error": "Meal Already Taken"})
		return
	} else if err != sql.ErrNoRows {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Database Error"})

		return
	}

	// Allow the meal
	_, err = database.DB.Exec("INSERT INTO meals (course_id, meal_type, date, time) VALUES (?, ?, ?, ?)", req.CourseID, mealType, today, currentTime.Format("15:04"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Meal allowed", "meal_type": mealType})
}

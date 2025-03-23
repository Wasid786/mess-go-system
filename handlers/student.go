package handlers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"messGo/database"
	"messGo/models"
	"net/http"
	"strings"
)

func CreateStudent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid Request Method", http.StatusMethodNotAllowed)

	}

	w.Header().Set("Content-Type", "application/json")

	var student models.Student
	err := json.NewDecoder(r.Body).Decode(&student)
	if err != nil {
		http.Error(w, "Invalid request Payload ", http.StatusBadRequest)
		return
	}

	query := "INSERT INTO students (enroll_no, name, hostel, course_id, registered_session, mess_slip) VALUES (?,?,?,?,?,?)"
	_, err = database.DB.Exec(query, student.EnrollNo, student.Name, student.Hostel, student.CourseID, student.RegisteredSession, student.MessSlip)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "No matcing records found ", http.StatusNotFound)
			log.Fatal(err)
			return
		}
		if strings.Contains(err.Error(), "Duplicate Entry") {
			http.Error(w, "Error: Student with this ID already Exists", http.StatusConflict)
			log.Println("Duplicate Entry error ", err)
			return
		}
		log.Println("Database execution error ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Student Created Successfully!"})

}

func GetStudents(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	rows, err := database.DB.Query("SELECT * FROM students")
	if err != nil {
		http.Error(w, `{"error": "Failed to fetch students"}`, http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var students []models.Student

	for rows.Next() {
		var student models.Student
		if err := rows.Scan(&student.ID, &student.EnrollNo, &student.Name, &student.Hostel, &student.CourseID, &student.RegisteredSession, &student.MessSlip); err != nil {
			http.Error(w, `{"error": "Failed to parse student data"}`, http.StatusInternalServerError)
			return
		}
		students = append(students, student)
	}

	data := map[string]interface{}{
		"Student": students,
	}

	// Render the template with data
	RenderTemplate(w, "checkstudents.tmpl", data)
}

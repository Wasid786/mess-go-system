package handlers

import (
	"encoding/json"
	"log"
	"messGo/database"
	"messGo/models"
	"net/http"
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

	query := "INSERT INTO students (id, enroll_no, name, hostel, course_id, registered_session, mess_slip) VALUES (?,?,?,?,?,?,?)"
	_, err = database.DB.Exec(query, student.ID, student.EnrollNo, student.Name, student.Hostel, student.CourseID, student.RegisteredSession, student.MessSlip)
	if err != nil {
		http.Error(w, "Failed to insert student", http.StatusInternalServerError)
		log.Fatal(err)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Student Created Successfully!"})

}

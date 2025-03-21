package models

type Student struct {
	ID                string `json:"id"`
	Name              string `json:"name"`
	Hostel            string `json:"hostel"`
	CourseID          string `json:"course_id"`
	EnrollNo          string `json:"enroll_no"`
	RegisteredSession string `json:"registered_session"`
	MessSlip          string `json:"mess_slip"`
}

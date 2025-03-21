package models

type Meal struct {
	ID       string `json:"id"`
	CourseID string `json:"course_id"`
	MealType string `json:"meal_type"`
	Date     string `json:"date"`
	Time     string `json:"time"`
}

package entity

import "time"

type Courses struct {
	ID          int       `gorm:"primaryKey" json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	VideoCourse string    `json:"video_course"`
	CategoryID  int       `json:"category_id"`
	MasteryID   int       `json:"mastery_id"`
	SubID       int       `json:"sub_type"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type CoursesInput struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	VideoCourse string `json:"video_course"`
	CategoryID  int    `json:"category_id"`
	MasteryID   int    `json:"mastery_id"`
	SubID       int    `json:"sub_type"`
}

type UpdateCoursesInput struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	VideoCourse string `json:"video_course"`
	CategoryID  int    `json:"category_id"`
	MasteryID   int    `json:"mastery_id"`
	SubID       int    `json:"sub_type"`
}

package entity

type Mastery struct {
	ID     int       `gorm:"primaryKey" json:"id"`
	Level  string    `json:"level"`
	Course []Courses `gorm:"foreignKey:MasteryID"`
}

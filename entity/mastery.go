package entity

type Mastery struct {
	ID     int       `gorm:"primaryKey" json:"id"`
	Level  string    `json:"level"`
	Course []Courses `gorm:"foreignKey:MasteryID"`
}

type MasteryInput struct {
	ID    int    `json:"id"`
	Level string `json:"level"`
}

type UpdateMasteryInput struct {
	Level string `json:"level"`
}

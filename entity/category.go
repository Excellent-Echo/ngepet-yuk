package entity

type Category struct {
	ID          int       `gorm:"primaryKey" json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Courses     []Courses `gorm:"foreignKey:CategoryID"`
}

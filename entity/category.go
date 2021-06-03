package entity

type Category struct {
	ID          int       `gorm:"primaryKey" json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Courses     []Courses `gorm:"foreignKey:CategoryID"`
}

type CategoryInput struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type UpdateCategoryInput struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

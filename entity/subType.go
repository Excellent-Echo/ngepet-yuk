package entity

type SubType struct {
	ID      int       `gorm:"primaryKey" json:"id"`
	SubType string    `json:"sub_type"`
	Price   int       `json:"price"`
	Period  int       `json:"period"`
	Course  []Courses `gorm:"foreignKey:SubID"`
}

type SubTypeInput struct {
	SubType string `json:"sub_type"`
	Price   int    `json:"price"`
	Period  int    `json:"period"`
}

type UpdateSubTypeInput struct {
	SubType string `json:"sub_type"`
	Price   int    `json:"price"`
	Period  int    `json:"period"`
}

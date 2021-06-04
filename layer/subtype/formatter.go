package subtype

import "ngepet-yuk/entity"

type SubtypeFormat struct {
	ID      int    `json:"id"`
	SubType string `json:"sub_type"`
	Price   int    `json:"price"`
	Period  int    `json:"period"`
}

func FormatSubtype(subtype entity.SubType) SubtypeFormat {
	var formatSubtype = SubtypeFormat{
		ID:      subtype.ID,
		SubType: subtype.SubType,
		Price:   subtype.Price,
		Period:  subtype.Period,
	}

	return formatSubtype
}

package user

import "ngepet-yuk/entity"

type UserFormat struct {
	ID       int    `json:"id"`
	UserName string `json:"username"`
	Email    string `json:"email"`
}

func FormatUser(user entity.User) UserFormat {
	var formatUser = UserFormat{
		ID:       user.ID,
		UserName: user.UserName,
		Email:    user.Email,
	}

	return formatUser
}

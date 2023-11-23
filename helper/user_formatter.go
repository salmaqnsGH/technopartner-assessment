package helper

import "technopartner/test/model/web"

type UserFormatter struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Token    string `json:"token"`
}

func FormatUser(user web.UserResponse, token string) UserFormatter {
	formatter := UserFormatter{
		ID:       user.ID,
		Username: user.Username,
		Token:    token,
	}
	return formatter
}

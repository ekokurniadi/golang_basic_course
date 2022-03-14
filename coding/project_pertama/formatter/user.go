package formatter

import "project_pertama/entity"

type UserFormatter struct {
	ID             int    `json:"id"`
	NamaLengkap    string `json:"nama_lengkap"`
	ProfilePicture string `json:"profile_picture"`
	JenisKelamin   string `json:"jenis_kelamin"`
	UserName       string `json:"user_name"`
	Password       string `json:"password"`
	Role           string `json:"role"`
}

func FormatUser(user entity.User) UserFormatter {
	userFormatter := UserFormatter{}
	userFormatter.ID = user.ID
	userFormatter.NamaLengkap = user.NamaLengkap
	userFormatter.ProfilePicture = user.ProfilePicture
	userFormatter.JenisKelamin = user.JenisKelamin
	userFormatter.UserName = user.UserName
	userFormatter.Password = user.Password
	userFormatter.Role = user.Role

	return userFormatter
}

func FormatUsers(users []entity.User) []UserFormatter {
	userFormatters := []UserFormatter{}
	for _, user := range users {
		userFormatter := FormatUser(user)
		userFormatters = append(userFormatters, userFormatter)
	}
	return userFormatters
}

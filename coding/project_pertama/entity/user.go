package entity

import "time"

type User struct {
	ID             int
	NamaLengkap    string
	JenisKelamin   string
	ProfilePicture string
	UserName       string
	Password       string
	Role           string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

package input

type UserInput struct {
	NamaLengkap    string `json:"nama_lengkap"`
	JenisKelamin   string `json:"jenis_kelamin"`
	ProfilePicture string `json:"profile_picture"`
	UserName       string `json:"user_name"`
	Password       string `json:"password"`
	Role           string `json:"role"`
}

type UserInputID struct {
	ID int `uri:"id"`
}

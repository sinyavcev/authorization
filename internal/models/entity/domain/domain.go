package domain

type User struct {
	ID       string
	Username string
}

type Password struct {
	ID           string
	PasswordHash string
	UserId       string
}

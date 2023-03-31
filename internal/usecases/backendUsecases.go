package usecases

type BackendUsecases interface {
	Signin()
	Signup()
	RefreshToken()
	Me()
	Logout()
}

type Authorization struct {
	BackendUsecases
}

func NewBackendUsecases() *Authorization {
	return &Authorization{}
}

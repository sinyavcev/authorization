package usecases

type backendUsecases interface {
	Signin()
	Signup()
	RefreshToken()
	Me()
	Logout()
}

type Authorization struct {
	backendUsecases
}

func NewBackendUsecases() *Authorization {
	return &Authorization{}
}

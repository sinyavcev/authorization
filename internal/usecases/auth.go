package usecases

import (
	"context"
	"fmt"
	"github.com/sinyavcev/authorization/internal/models/entity/backendModels"
	"github.com/sinyavcev/authorization/internal/models/entity/domain"
	"golang.org/x/crypto/bcrypt"
)

func (u *Usecases) SignUp(data *backendModels.SignUpRequest) (*domain.User, error) {
	user, err := u.repository.CreateUser(
		context.Background(),
		domain.User{Username: data.Username},
		domain.Password{PasswordHash: HashPassword(data.Password)},
	)
	if err != nil {
		return nil, fmt.Errorf("repository.CreateUser: %w", err)
	}
	return &user, nil
}
func (u *Usecases) SignIn(data *backendModels.SignInRequest) (*domain.User, error) {
	user, password, err := u.repository.GetUser(
		context.Background(),
		domain.User{Username: data.Username},
	)
	if err != nil {
		fmt.Errorf("GetUser: %w", err)
		return nil, err
	}
	err = VerifyPassword(password.PasswordHash, data.Password)
	if err != nil {
		fmt.Errorf("VerifyPassword: %w", err)
		return nil, err
	}
	return &user, nil
}

func HashPassword(password string) string {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPassword)
}

func VerifyPassword(hashedPassword string, candidatePassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(candidatePassword))
}

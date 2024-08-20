package usecases

import (
	"context"
	"github.com/sinyavcev/authorization/internal/models/entity/domain"
)

type Repository interface {
	CreateUser(ctx context.Context, user domain.User, password domain.Password) (domain.User, error)
	GetUser(ctx context.Context, user domain.User) (domain.User, domain.Password, error)
}

type Usecases struct {
	repository Repository
}

func NewUsecases(repository Repository) *Usecases {
	return &Usecases{
		repository: repository,
	}
}

package service

type Authorization interface {
}

type Repository struct {
	Authorization
}

func NewRepository() *Repository {
	return &Repository{}
}

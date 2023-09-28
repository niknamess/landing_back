package service

import "sober_driver/internal/repository"

type Service struct {
	UserAction IUserAction
	FormAction IFormAction
}

func NewService(deps *Deps) *Service {
	return &Service{
		UserAction: newUserAction(deps.Repo),
		FormAction: newFormAction(deps.Repo),
	}
}

type Deps struct {
	Repo *repository.Repository
}

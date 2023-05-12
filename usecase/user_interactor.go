package usecase

import (
	"github.com/Luftalian/Clean_Architecture/domain"
)

type UserInteractor struct {
	UserRepository UserRepository
}

func (interactor *UserInteractor) Add(u domain.User) (domain.User, error) {
	user, err := interactor.UserRepository.Store(u)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func (interactor *UserInteractor) Users() (domain.Users, error) {
	users, err := interactor.UserRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (interactor *UserInteractor) UserByID(id domain.UUID) (domain.User, error) {
	user, err := interactor.UserRepository.FindByID(id)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}

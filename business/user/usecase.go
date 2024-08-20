package users

import (
	"arifptmx/helpers"
	"errors"
)

type UserUseCase struct {
	repo UserRepoInterface
}

func NewUseCase(userRepo UserRepoInterface) UserUseCaseInterface {
	return &UserUseCase{
		repo: userRepo,
	}
}

func (userUseCase *UserUseCase) SignUp(user User) (User, error) {
	// check request
	if user.Name == "" {
		return User{}, errors.New("name empty")
	}

	if user.Email == "" {
		return User{}, errors.New("email empty")
	}

	if user.Password == "" {
		return User{}, errors.New("password empty")
	}

	// hash password
	hash, _ := helpers.HashPassword(user.Password)

	user.Password = hash

	userRepo, err := userUseCase.repo.SignUp(user)

	if err != nil {
		return User{}, err
	}

	return userRepo, nil
}

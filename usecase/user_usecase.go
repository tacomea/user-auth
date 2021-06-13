package usecase

import "userCreation/domain"

type userUsecase struct {
	userRepo domain.UserRepository
}

func NewUserUsecase(ur domain.UserRepository) domain.UserUseCase {
	return &userUsecase{
		userRepo: ur,
	}
}

func (u *userUsecase) Create(user domain.User) error {
	err := u.userRepo.Create(user)
	if err != nil {
		return err
	}
	return nil
}

func (u *userUsecase) Delete(email string) error {
	err := u.userRepo.Delete(email)
	if err != nil {
		return err
	}
	return nil
}

func (u *userUsecase) Check(email string) (domain.User, error) {
	user, err := u.userRepo.Check(email)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}

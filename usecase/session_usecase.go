package usecase

import "userCreation/domain"

type sessionUsecase struct {
	sessionRepo domain.SessionRepository
}

func NewSessionUsecase(sr domain.SessionRepository) domain.SessionUseCase {
	return &sessionUsecase{
		sessionRepo: sr,
	}
}

func (u *sessionUsecase) Store(session domain.Session) error {
	err := u.sessionRepo.Store(session)
	if err != nil {
		return err
	}
	return nil
}

func (u *sessionUsecase) Delete(id string) error {
	err := u.sessionRepo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

func (u *sessionUsecase) Load(id string) (domain.Session, error) {
	user, err := u.sessionRepo.Load(id)
	if err != nil {
		return domain.Session{}, err
	}
	return user, nil
}

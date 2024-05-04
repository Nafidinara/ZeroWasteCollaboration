package usecases

import (
	"context"
	"time"

	"redoocehub/domains"
)

type profileUsecase struct {
	userRepository domains.UserRepository
	contextTimeout time.Duration
}

func NewProfileUsecase(userRepository domains.UserRepository, timeout time.Duration) domains.ProfileUsecase {
	return &profileUsecase{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}

func (pu *profileUsecase) GetProfileByID(c context.Context, userID string) (*domains.Profile, error) {
	user, err := pu.userRepository.GetByID(userID)
	if err != nil {
		return nil, err
	}

	return &domains.Profile{Name: user.Name, Email: user.Email}, nil
}

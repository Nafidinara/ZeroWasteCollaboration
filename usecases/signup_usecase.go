package usecases

import (
	"context"
	"time"

	"redoocehub/domains"
	"redoocehub/internal/tokenutil"
)

type signupUsecase struct {
	userRepository domains.UserRepository
	contextTimeout time.Duration
}

func NewSignupUsecase(userRepository domains.UserRepository, timeout time.Duration) domains.SignupUsecase {
	return &signupUsecase{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}

func (su *signupUsecase) Create(c context.Context, user *domains.User) error {
	return su.userRepository.Create(user)
}

func (su *signupUsecase) GetUserByEmail(c context.Context, email string) (domains.User, error) {
	return su.userRepository.GetByEmail(email)
}

func (su *signupUsecase) CreateAccessToken(user *domains.User, secret string, expiry int) (accessToken string, err error) {
	return tokenutil.CreateAccessToken(user, secret, expiry)
}

func (su *signupUsecase) CreateRefreshToken(user *domains.User, secret string, expiry int) (refreshToken string, err error) {
	return tokenutil.CreateRefreshToken(user, secret, expiry)
}

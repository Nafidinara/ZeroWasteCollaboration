package usecases

import (
	"context"
	"time"

	"redoocehub/domains"
	"redoocehub/internal/tokenutil"
)

type loginUsecase struct {
	userRepository domains.UserRepository
	contextTimeout time.Duration
}

func NewLoginUsecase(userRepository domains.UserRepository, timeout time.Duration) domains.LoginUsecase {
	return &loginUsecase{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}

func (lu *loginUsecase) GetUserByEmail(c context.Context, email string) (domains.User, error) {
	return lu.userRepository.GetByEmail(email)
}

func (lu *loginUsecase) CreateAccessToken(user *domains.User, secret string, expiry int) (accessToken string, err error) {
	return tokenutil.CreateAccessToken(user, secret, expiry)
}

func (lu *loginUsecase) CreateRefreshToken(user *domains.User, secret string, expiry int) (refreshToken string, err error) {
	return tokenutil.CreateRefreshToken(user, secret, expiry)
}

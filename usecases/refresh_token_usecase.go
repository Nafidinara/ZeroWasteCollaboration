package usecases

import (
	"context"
	"time"

	"redoocehub/domains"
	"redoocehub/internal/tokenutil"
)

type refreshTokenUsecase struct {
	userRepository domains.UserRepository
	contextTimeout time.Duration
}

func NewRefreshTokenUsecase(userRepository domains.UserRepository, timeout time.Duration) domains.RefreshTokenUsecase {
	return &refreshTokenUsecase{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}

func (ru *refreshTokenUsecase) GetUserByID(c context.Context, id string) (domains.User, error) {
	return ru.userRepository.GetByID(id)
}

func (ru *refreshTokenUsecase) CreateAccessToken(user *domains.User, secret string, expiry int) (accessToken string, err error) {
	return tokenutil.CreateAccessToken(user, secret, expiry)
}

func (ru *refreshTokenUsecase) CreateRefreshToken(user *domains.User, secret string, expiry int) (refreshToken string, err error) {
	return tokenutil.CreateRefreshToken(user, secret, expiry)
}

func (ru *refreshTokenUsecase) ExtractIDFromToken(requestToken string, secret string) (string, error) {
	return tokenutil.ExtractIDFromToken(requestToken, secret)
}
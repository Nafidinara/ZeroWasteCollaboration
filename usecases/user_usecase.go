package usecases

import (
	"context"
	"time"

	"redoocehub/domains/entities"
	"redoocehub/internal/tokenutil"
)

type userUsecase struct {
	userRepository entities.UserRepository
	contextTimeout time.Duration
}

func NewUserUsecase(userRepository entities.UserRepository, timeout time.Duration) entities.UserUsecase {
	return &userUsecase{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}

func (u *userUsecase) Create(c context.Context, user *entities.User) error {
	return u.userRepository.Create(user)
}

func (u *userUsecase) GetUserByEmail(c context.Context, email string) (entities.User, error) {
	return u.userRepository.GetByEmail(email)
}

func (u *userUsecase) GetUserByID(c context.Context, id string) (entities.User, error) {
	return u.userRepository.GetByID(id)
}

func (u *userUsecase) GetProfileByID(c context.Context, userID string) (*entities.User, error) {
	user, err := u.userRepository.GetByID(userID)
	if err != nil {
		return nil, err
	}

	return &entities.User{
		ID:           user.ID,
		Username:     user.Username,
		ProfileImage: user.ProfileImage,
		Email:        user.Email,
		FullName:     user.FullName,
		Gender:       user.Gender,
		UpdatedAt:    user.UpdatedAt,
		CreatedAt:    user.CreatedAt,
		DeletedAt:    user.DeletedAt,
	}, nil
}

func (u *userUsecase) CreateAccessToken(user *entities.User, secret string, expiry int) (accessToken string, err error) {
	return tokenutil.CreateAccessToken(user, secret, expiry)
}

func (u *userUsecase) CreateRefreshToken(user *entities.User, secret string, expiry int) (refreshToken string, err error) {
	return tokenutil.CreateRefreshToken(user, secret, expiry)
}

func (u *userUsecase) ExtractIDFromToken(requestToken string, secret string) (string, error) {
	return tokenutil.ExtractIDFromToken(requestToken, secret)
}

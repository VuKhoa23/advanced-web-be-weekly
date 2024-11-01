package serviceimplement

import (
	"context"
	"errors"

	"github.com/VuKhoa23/advanced-web-be/internal/domain/entity"
	"github.com/VuKhoa23/advanced-web-be/internal/domain/model"
	"github.com/VuKhoa23/advanced-web-be/internal/repository"
	"github.com/VuKhoa23/advanced-web-be/internal/service"
	utils "github.com/VuKhoa23/advanced-web-be/internal/utils/jwt_utils"
	stringutils "github.com/VuKhoa23/advanced-web-be/internal/utils/string_utils"
)

type UserService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) service.UserService {
	return &UserService{userRepository: userRepository}
}

// Register implements service.UserService.
func (service *UserService) Register(ctx context.Context, userRequest model.UserRequest) (*entity.User, error) {
	user := &entity.User{
		Username: userRequest.Username,
		Password: stringutils.HashPassword(userRequest.Password),
	}

	err := service.userRepository.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// Login implements service.UserService.
func (service *UserService) Login(ctx context.Context, userRequest model.UserRequest) (string, error) {
	// get user by username
	user, err := service.userRepository.GetUserByUsername(ctx, userRequest.Username)
	if err != nil {
		return "", err
	}

	// compare and hash password
	if !stringutils.CompareHashAndPassword(user.Password, userRequest.Password) {
		// If the password is incorrect, return an error
		return "", errors.New("invalid username or password")
	}

	// generate token
	token, err := utils.CreateToken(user.Id, user.Username)
	if err != nil {
		return "", err
	}

	return token, nil
}
package serviceimplement

import (
	"context"
	"github.com/VuKhoa23/advanced-web-be/internal/domain/entity"
	"github.com/VuKhoa23/advanced-web-be/internal/domain/model"
	"github.com/VuKhoa23/advanced-web-be/internal/repository"
	"github.com/VuKhoa23/advanced-web-be/internal/service"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) service.UserService {
	return &UserService{userRepository: userRepository}
}

func (service *UserService) Login(ctx context.Context, loginRequest model.LoginRequest) (*entity.User, error) {
	user, err := service.userRepository.FindUserByUsername(ctx, loginRequest.Username)
	if err != nil {
		return nil, err
	}
	//check password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password))
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (service *UserService) Register(ctx context.Context, register model.RegisterRequest) (string, error) {
	//hash password before save
	hashPW, err := bcrypt.GenerateFromPassword([]byte(register.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	user := &entity.User{
		Username: register.Username,
		Password: string(hashPW),
	}
	err = service.userRepository.CreateUser(ctx, user)
	if err != nil {
		return "", err
	}
	return user.Username, nil
}

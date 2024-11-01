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

func (service *UserService) CheckPassword(ctx context.Context, userRequest model.UserRequest) (*entity.User, error) {
	user, err := service.userRepository.FindUserByUsername(ctx, userRequest.UserName)
	if err != nil {
		return nil, err
	}
	//check password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userRequest.Password))
	if err != nil {
		return nil, err
	}
	return user, nil
}
func (service *UserService) CreateUser(ctx context.Context, userRequest model.UserRequest) (string, error) {
	//hash password before save
	hashPW, err := bcrypt.GenerateFromPassword([]byte(userRequest.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	user := &entity.User{
		UserName: userRequest.UserName,
		Password: string(hashPW),
	}
	err = service.userRepository.CreateUser(ctx, user)
	if err != nil {
		return "", err
	}
	return user.UserName, nil
}

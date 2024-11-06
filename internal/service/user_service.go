package service

import (
	"context"
	"github.com/VuKhoa23/advanced-web-be/internal/domain/entity"
	"github.com/VuKhoa23/advanced-web-be/internal/domain/model"
)

type UserService interface {
	Login(ctx context.Context, loginRequest model.LoginRequest) (*entity.User, error)
	Register(ctx context.Context, registerRequest model.RegisterRequest) (string, error)
}

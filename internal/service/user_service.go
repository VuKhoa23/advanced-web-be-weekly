package service

import (
	"context"

	"github.com/VuKhoa23/advanced-web-be/internal/domain/entity"
	"github.com/VuKhoa23/advanced-web-be/internal/domain/model"
)

type UserService interface {
	Register(c context.Context, userRequest model.UserRequest) (*entity.User, error)
	Login(c context.Context, userRequest model.UserRequest) (string, error)
}
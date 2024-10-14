package service

import (
	"context"
	"github.com/VuKhoa23/advanced-web-be/internal/domain/entity"
)

type ActorService interface {
	GetAllActor(ctx context.Context) []entity.Actor
	GetActorById(ctx context.Context, id int64) (*entity.Actor, error)
	CreateActor(ctx context.Context, actor *entity.Actor) error
	UpdateActor(ctx context.Context, actor *entity.Actor) error
	DeleteActor(ctx context.Context, id int64) error
}

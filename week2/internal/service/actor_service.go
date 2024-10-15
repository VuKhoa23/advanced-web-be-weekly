package service

import (
	"context"
	"github.com/VuKhoa23/advanced-web-be/internal/domain/entity"
	"github.com/VuKhoa23/advanced-web-be/internal/domain/model"
)

type ActorService interface {
	GetAllActor(ctx context.Context) []entity.Actor
	GetActorById(ctx context.Context, id int64) (*entity.Actor, error)
	CreateActor(ctx context.Context, actorRequest model.ActorRequest) (*entity.Actor, error)
	UpdateActor(ctx context.Context, actorRequest model.ActorRequest, actorId int64) (*entity.Actor, error)
	DeleteActor(ctx context.Context, id int64) error
}

package repository

import (
	"context"
	"github.com/VuKhoa23/advanced-web-be/internal/domain/entity"
)

type ActorRepository interface {
	GetAllActors(c context.Context) []entity.Actor
	GetActorByID(c context.Context, id int64) (*entity.Actor, error)
	CreateActor(c context.Context, actor *entity.Actor) error
	UpdateActor(c context.Context, actor *entity.Actor, actorId int64) (*entity.Actor, error)
	DeleteActor(c context.Context, id int64) error
}

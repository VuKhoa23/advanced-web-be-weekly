package serviceimplement

import (
	"context"
	"github.com/VuKhoa23/advanced-web-be/internal/domain/entity"
	"github.com/VuKhoa23/advanced-web-be/internal/domain/model"
	"github.com/VuKhoa23/advanced-web-be/internal/repository"
	"github.com/VuKhoa23/advanced-web-be/internal/service"
)

type ActorService struct {
	actorRepository repository.ActorRepository
}

func NewActorService(actorRepository repository.ActorRepository) service.ActorService {
	return &ActorService{actorRepository: actorRepository}
}

func (service *ActorService) GetAllActor(ctx context.Context) []entity.Actor {
	return service.actorRepository.GetAllActors(ctx)
}

func (service *ActorService) GetActorById(ctx context.Context, id int64) (*entity.Actor, error) {
	return service.actorRepository.GetActorByID(ctx, id)
}

func (service *ActorService) CreateActor(ctx context.Context, actorRequest model.ActorRequest) (*entity.Actor, error) {
	actor := &entity.Actor{FirstName: actorRequest.FirstName, LastName: actorRequest.LastName}
	err := service.actorRepository.CreateActor(ctx, actor)
	if err != nil {
		return nil, err
	}
	return actor, nil
}

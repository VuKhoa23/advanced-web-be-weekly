package repositoryimplement

import (
	"context"
	"github.com/VuKhoa23/advanced-web-be/internal/database"
	"github.com/VuKhoa23/advanced-web-be/internal/domain/entity"
	"github.com/VuKhoa23/advanced-web-be/internal/repository"
	"gorm.io/gorm"
)

type ActorRepository struct {
	db *gorm.DB
}

func NewActorRepository(db database.Db) repository.ActorRepository {
	return &ActorRepository{db: db}
}

func (repo *ActorRepository) GetAllActors(ctx context.Context) []entity.Actor {
	var actors []entity.Actor
	result := repo.db.WithContext(ctx).Find(&actors)
	if result.Error != nil {
		return []entity.Actor{}
	}
	return actors
}

func (repo *ActorRepository) GetActorByID(ctx context.Context, id int64) (*entity.Actor, error) {
	var actor entity.Actor
	result := repo.db.WithContext(ctx).First(&actor, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &actor, nil
}

func (repo *ActorRepository) CreateActor(ctx context.Context, actor *entity.Actor) error {
	err := repo.db.WithContext(ctx).Create(actor).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *ActorRepository) UpdateActor(ctx context.Context, actor *entity.Actor) error {
	//update all columns of an actor
	if err := repo.db.WithContext(ctx).Where("actor_id = ?", actor.ID).Save(actor).Error; err != nil {
		//maybe dont need
		//if errors.Is(err, gorm.ErrRecordNotFound) {
		//	return nil, err
		//}
		return err
	}
	return nil
}

func (repo *ActorRepository) DeleteActor(ctx context.Context, id int64) error {
	return repo.db.Transaction(func(tx *gorm.DB) error {
		//check actor exists
		var actor entity.Actor
		if err := tx.WithContext(ctx).First(&actor, id).Error; err != nil {
			return err
		}

		// Delete from film_actor
		if err := tx.Exec(`
            DELETE fa FROM film_actor fa
            WHERE fa.actor_id = ?`, id).Error; err != nil {
			return err
		}

		//Delete actor
		if err := tx.Delete(&actor).Error; err != nil {
			return err
		}
		return nil
	})
}

package repositoryimplement

import (
	"context"

	"github.com/VuKhoa23/advanced-web-be/internal/database"
	"github.com/VuKhoa23/advanced-web-be/internal/domain/entity"
	"github.com/VuKhoa23/advanced-web-be/internal/repository"
	"gorm.io/gorm"
)

type FilmRepository struct {
	db *gorm.DB
}

func NewFilmRepository(db database.Db) repository.FilmRepository {
	return &FilmRepository{db: db}
}

func (repo *FilmRepository) GetFilmByID(ctx context.Context, id int64) (*entity.Film, error) {
	var film entity.Film
	result := repo.db.WithContext(ctx).First(&film, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &film, nil
}

func (repo *FilmRepository) DeleteFilm(ctx context.Context, id int64) error {
	return repo.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// Delete from film_actor
		if err := tx.Exec(`
            DELETE fa FROM film_actor fa
            JOIN film f ON f.film_id = fa.film_id
            WHERE f.film_id = ?`, id).Error; err != nil {
			return err
		}

		// Delete from film_category
		if err := tx.Exec(`
            DELETE fc FROM film_category fc
            JOIN film f ON f.film_id = fc.film_id
            WHERE f.film_id = ?`, id).Error; err != nil {
			return err
		}

		// Delete from payment (inventory -> rental -> payment)
		if err := tx.Exec(`
            DELETE p FROM payment p
            JOIN rental r ON r.rental_id = p.rental_id
            JOIN inventory i ON i.inventory_id = r.inventory_id
            JOIN film f ON f.film_id = i.film_id
            WHERE f.film_id = ?`, id).Error; err != nil {
			return err
		}

		// Delete from rental (inventory -> rental)
		if err := tx.Exec(`
            DELETE r FROM rental r
            JOIN inventory i ON i.inventory_id = r.inventory_id
            JOIN film f ON f.film_id = i.film_id
            WHERE f.film_id = ?`, id).Error; err != nil {
			return err
		}

		// Delete from inventory
		if err := tx.Exec(`
            DELETE i FROM inventory i
            JOIN film f ON f.film_id = i.film_id
            WHERE f.film_id = ?`, id).Error; err != nil {
			return err
		}

		// Delete film
		if err := tx.Exec(`
            DELETE FROM film
            WHERE film_id = ?`, id).Error; err != nil {
			return err
		}

		return nil
	})
}

package repositories

import (
	"log"

	"github.com/dusk-chancellor/mego-favorite/internal/models"
	"github.com/jmoiron/sqlx"
)

const (
	element = "favorite_repository"
)

type FavoriteRepository interface {
	Create(favorite *models.Favorite) (string, error)
	Increment(postID string) (string, error)
	Decrement(postID string) (string, error)
}

type favoriteRepository struct {
	db *sqlx.DB
}

func NewFavoriteRepository(db *sqlx.DB) FavoriteRepository {
	return &favoriteRepository{
		db: db,
	}
}

func (r *favoriteRepository) Create(favorite *models.Favorite) (string, error) {
	q := `INSERT INTO favorites (post_id, count) VALUES ($1, $2)`

	_, err := r.db.Exec(q, favorite.PostID, favorite.Count)
	if err != nil {
		log.Printf("Element: %s | Failed to create favorite: %v", element, err)
		return "", err
	}
	return favorite.PostID, nil
}

func (r *favoriteRepository) Increment(postID string) (string, error) {
	q := `UPDATE favorites SET count = count + 1 WHERE post_id = $1`

	_, err := r.db.Exec(q, postID)
	if err != nil {
		log.Printf("Element: %s | Failed to increment favorite: %v", element, err)
		return "", err
	}
	return postID, nil
}

func (r *favoriteRepository) Decrement(postID string) (string, error) {
	q := `UPDATE favorites SET count = count - 1 WHERE post_id = $1`

	_, err := r.db.Exec(q, postID)
	if err != nil {
		log.Printf("Element: %s | Failed to decrement favorite: %v", element, err)
		return "", err
	}
	return postID, nil
}



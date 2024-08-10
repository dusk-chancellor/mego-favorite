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
	Exists(favorite models.Favorite) bool
	Add(favorite models.Favorite) (string, string, error)
	Delete(favorite models.Favorite) (string, string, error)
	Count(postID string) int32
}

type favoriteRepository struct {
	db *sqlx.DB
}

func NewFavoriteRepository(db *sqlx.DB) FavoriteRepository {
	return &favoriteRepository{
		db: db,
	}
}

func (r *favoriteRepository) Add(favorite models.Favorite) (string, string, error) {
	q := `INSERT INTO favorites (user_id, post_id) VALUES ($1, $2) RETURNING user_id, post_id;`

	var userId, postId string
	err := r.db.QueryRow(q, favorite.UserId, favorite.PostID).Scan(&userId, &postId)
	if err != nil {
		log.Printf("Element: %s | Failed to add favorite: %v", element, err)
		return "", "", err
	}
	return userId, postId, nil
}

func (r *favoriteRepository) Exists(favorite models.Favorite) bool {
	q := `SELECT EXISTS(SELECT 1 FROM favorites WHERE user_id = $1 AND post_id = $2);`

	var exists bool
	err := r.db.QueryRow(q, favorite.UserId, favorite.PostID).Scan(&exists)
	if err != nil {
		log.Printf("Element: %s | Failed to check if favorite exists: %v", element, err)
		return false
	}
	return exists
}

func (r *favoriteRepository) Delete(favorite models.Favorite) (string, string, error) {
	q := `DELETE FROM favorites WHERE user_id = $1 AND post_id = $2 RETURNING user_id, post_id;`

	var userId, postId string
	err := r.db.QueryRow(q, favorite.UserId, favorite.PostID).Scan(&userId, &postId)
	if err != nil {
		log.Printf("Element: %s | Failed to delete favorite: %v", element, err)
		return "", "", err
	}
	return userId, postId, nil
}

func (r *favoriteRepository) Count(postID string) int32 {
	q := `SELECT COUNT(*) FROM favorites WHERE post_id = $1;`

	var count int32
	err := r.db.QueryRow(q, postID).Scan(&count)
	if err != nil {
		log.Printf("Element: %s | Failed to count favorites: %v", element, err)
		return 0
	}
	return count
}

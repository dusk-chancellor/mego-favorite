package repositories

import (
	"github.com/dusk-chancellor/mego-favorite/internal/models"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
)

const (
	element = "favorite_repository"
)

type FavoriteRepository interface {
	Exists(favorite models.Favorite) (bool, error)
	Add(favorite models.Favorite) (string, string, error)
	Delete(favorite models.Favorite) (string, string, error)
	Find(startIndex, pageSize int) ([]*models.Favorite, error)
	Count(postID string) (int32, error)
}

type favoriteRepository struct {
	db *sqlx.DB
	redis *redis.Client
}

func NewFavoriteRepository(db *sqlx.DB, rdb *redis.Client) FavoriteRepository {
	return &favoriteRepository{
		db: db,
		redis: rdb,
	}
}
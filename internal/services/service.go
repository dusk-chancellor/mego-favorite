package services

import (
	"github.com/dusk-chancellor/mego-favorite/internal/models"
	"github.com/dusk-chancellor/mego-favorite/internal/repositories"
)

const (
	element = "favorite_service"
)

type FavoriteService interface {
	Exists(favorite models.Favorite) (bool, error)
	Add(favorite models.Favorite) (string, string, error)
	Delete(favorite models.Favorite) (string, string, error)
	Find(pageSize int, pageToken string) ([]*models.Favorite, string, error)
	Count(postID string) (int32, error)
}

type favoriteService struct {
	favoriteRepository repositories.FavoriteRepository
}

func NewFavoriteService(favoriteRepo repositories.FavoriteRepository) FavoriteService {
	return &favoriteService{
		favoriteRepository: favoriteRepo,
	}
}

package services

import (
	"log"

	"github.com/dusk-chancellor/mego-favorite/internal/models"
	"github.com/dusk-chancellor/mego-favorite/internal/repositories"
)

const (
	element = "favorite_service"
)

type FavoriteService interface {
	Exists(favorite models.Favorite) bool
	Add(favorite models.Favorite) (string, string, error)
	Delete(favorite models.Favorite) (string, string, error)
	Count(postID string) int32
}

type favoriteService struct {
	favoriteRepository repositories.FavoriteRepository
	favoriteLocalCache *favoriteLocalCache
}

func NewFavoriteService(favoriteRepo repositories.FavoriteRepository, favoriteLocalCache *favoriteLocalCache) FavoriteService {
	return &favoriteService{
		favoriteRepository: favoriteRepo,
		favoriteLocalCache: favoriteLocalCache,
	}
}

func (s *favoriteService) Exists(favorite models.Favorite) bool {
	return s.favoriteLocalCache.Exists(favorite)
}

func (s *favoriteService) Add(favorite models.Favorite) (string, string, error) {
	userId, postId, err := s.favoriteRepository.Add(favorite)
	if err != nil {
		log.Printf("Element: %s | Failed to add favorite: %v", element, err)
		return "", "", err
	}
	go s.favoriteLocalCache.Add(favorite)
	return userId, postId, nil
}

func (s *favoriteService) Delete(favorite models.Favorite) (string, string, error) {
	userId, postId, err := s.favoriteRepository.Delete(favorite)
	if err != nil {
		log.Printf("Element: %s | Failed to delete favorite: %v", element, err)
		return "", "", err
	}
	go s.favoriteLocalCache.Delete(favorite)
	return userId, postId, nil
}

func (s *favoriteService) Count(postID string) int32 {
	return s.favoriteLocalCache.Count(postID)
}


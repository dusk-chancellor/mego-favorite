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
}

func NewFavoriteService(favoriteRepo repositories.FavoriteRepository) FavoriteService {
	return &favoriteService{
		favoriteRepository: favoriteRepo,
	}
}

func (s *favoriteService) Exists(favorite models.Favorite) bool {
	return s.favoriteRepository.Exists(favorite)
}

func (s *favoriteService) Add(favorite models.Favorite) (string, string, error) {
	if s.Exists(favorite) {
		return "", "", nil
	}
	userId, postId, err := s.favoriteRepository.Add(favorite)
	if err != nil {
		log.Printf("Element: %s | Failed to add favorite: %v", element, err)
		return "", "", err
	}
	return userId, postId, nil
}

func (s *favoriteService) Delete(favorite models.Favorite) (string, string, error) {
	userId, postId, err := s.favoriteRepository.Delete(favorite)
	if err != nil {
		log.Printf("Element: %s | Failed to delete favorite: %v", element, err)
		return "", "", err
	}
	return userId, postId, nil
}

func (s *favoriteService) Count(postID string) int32 {
	return s.favoriteRepository.Count(postID)
}


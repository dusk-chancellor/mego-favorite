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
	Exists(postID string) bool
	Add(postID string) (string, error)
	Delete(postID string) (string, error)
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

func (s *favoriteService) Exists(postID string) bool {
	return s.favoriteLocalCache.Exists(postID)
}

func (s *favoriteService) Add(postID string) (string, error) {
	if exists := s.favoriteLocalCache.Exists(postID); exists {
		id, err := s.favoriteRepository.Increment(postID)
		if err != nil {
			log.Printf("Element: %s | Failed to add favorite: %v", element, err)
			return "", err
		}
		go s.favoriteLocalCache.Increment(postID)
		return id, nil
	}
	favorite := &models.Favorite{PostID: postID, Count: 1}
	id, err := s.favoriteRepository.Create(favorite)
	if err != nil {
		log.Printf("Element: %s | Failed to add favorite: %v", element, err)
		return "", err
	}
	go s.favoriteLocalCache.Add(postID)
	return id, nil
}

func (s *favoriteService) Delete(postID string) (string, error) {
	id, err := s.favoriteRepository.Decrement(postID)
	if err != nil {
		log.Printf("Element: %s | Failed to delete favorite: %v", element, err)
		return "", err
	}
	go s.favoriteLocalCache.Decrement(postID)
	return id, nil
}

func (s *favoriteService) Count(postID string) int32 {
	return s.favoriteLocalCache.Count(postID)
}


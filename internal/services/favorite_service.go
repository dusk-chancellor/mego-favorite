package services

import (
	"log"

	"github.com/dusk-chancellor/mego-favorite/internal/models"
	"github.com/dusk-chancellor/mego-favorite/pkg/utils"
)

func (s *favoriteService) Exists(favorite models.Favorite) (bool, error) {
	return s.favoriteRepository.Exists(favorite)
}

func (s *favoriteService) Add(favorite models.Favorite) (string, string, error) {
	exists, _ := s.Exists(favorite)
	if exists {
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

func (s *favoriteService) Find(pageSize int, pageToken string) ([]*models.Favorite, string, error) {
	var err error
	if pageSize < 1 {
		pageSize = 10
	}

	startIndex := 0
	if pageToken != "" {
		startIndex, err = utils.DecodePageToken(pageToken)
		if err != nil {
			log.Printf("Element: %s | Failed to decode page token: %v", element, err)
			return nil, "", err
		}
	}

	favorites, err := s.favoriteRepository.Find(startIndex, pageSize+1)
	if err != nil {
		log.Printf("Element: %s | Failed to find favorites: %v", element, err)
		return nil, "", err
	}

	var nextPageToken string
	if len(favorites) > pageSize {
		nextPageToken = utils.EncodePageToken(startIndex + pageSize)
		favorites = favorites[:pageSize]
	}

	return favorites, nextPageToken, nil
}

func (s *favoriteService) Count(postID string) (int32, error) {
	return s.favoriteRepository.Count(postID)
}


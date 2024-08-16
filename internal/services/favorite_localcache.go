package services

import (
	"sync"

	"github.com/dusk-chancellor/mego-favorite/internal/models"
)

type favoriteLocalCache struct {
	favorites map[string]string
	sync.Mutex
}

func NewFavoriteLocalCache() *favoriteLocalCache {
	return &favoriteLocalCache{
		favorites: make(map[string]string),
	}
}

func (lc *favoriteLocalCache) Add(favorite models.Favorite) {
	lc.Lock()
	defer lc.Unlock()
	lc.favorites[favorite.PostID] = favorite.UserId
}

func (lc *favoriteLocalCache) Delete(favorite models.Favorite) {
	lc.Lock()
	defer lc.Unlock()
	delete(lc.favorites, favorite.PostID)
}

func (lc *favoriteLocalCache) Exists(favorite models.Favorite) bool {
	_, exists := lc.favorites[favorite.PostID]
	return exists
}

func (lc *favoriteLocalCache) Count(postID string) int32 {
	count := len(lc.favorites)
	return int32(count)
}

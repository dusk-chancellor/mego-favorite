package services

import "sync"

type favoriteLocalCache struct {
	posts map[string]int32
	sync.Mutex
}

func NewFavoriteLocalCache() *favoriteLocalCache {
	return &favoriteLocalCache{}
}

func (lc *favoriteLocalCache) Add(postID string) {
	lc.Lock()
	defer lc.Unlock()
	lc.posts[postID] = 1
}

func (lc *favoriteLocalCache) Exists(postID string) bool {
	_, ok := lc.posts[postID]
	return ok
}

func (lc *favoriteLocalCache) Increment(postID string) {
	lc.Lock()
	defer lc.Unlock()
	lc.posts[postID] += 1
}

func (lc *favoriteLocalCache) Decrement(postID string) {
	lc.Lock()
	defer lc.Unlock()
	lc.posts[postID] -= 1
}

func (lc *favoriteLocalCache) Count(postID string) int32 {
	return lc.posts[postID]
}

package repositories

// redis: key structure -> user_id:post_id

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/dusk-chancellor/mego-favorite/internal/models"
)

func (r *favoriteRepository) Add(favorite models.Favorite) (string, string, error) {
	q := `INSERT INTO favorites (user_id, post_id) VALUES ($1, $2) RETURNING user_id, post_id;`

	var userId, postId string
	err := r.db.QueryRow(q, favorite.UserId, favorite.PostId).Scan(&userId, &postId)
	if err != nil {
		log.Printf("Element: %s | Failed to add favorite in db: %v", element, err)
		return "", "", err
	}

	key := fmt.Sprintf("%s:%s", favorite.UserId, favorite.PostId)
	_, err = r.redis.Set(context.Background(), key, 1, 24*time.Hour).Result()
	if err != nil {
		log.Printf("Element: %s | Failed to add favorite in redis: %v", element, err)
	}

	return userId, postId, nil
}

func (r *favoriteRepository) Exists(favorite models.Favorite) (bool, error) {
	key := fmt.Sprintf("%s:%s", favorite.UserId, favorite.PostId)
	exists, _ := r.redis.Get(context.Background(), key).Bool()
	if exists {
		return true, nil
	}

	q := `SELECT EXISTS(SELECT 1 FROM favorites WHERE user_id = $1 AND post_id = $2);`

	err := r.db.QueryRow(q, favorite.UserId, favorite.PostId).Scan(&exists)
	if err != nil {
		log.Printf("Element: %s | Failed to check if favorite exists in db: %v", element, err)
		return false, err
	}
	return exists, nil
}

func (r *favoriteRepository) Delete(favorite models.Favorite) (string, string, error) {
	q := `DELETE FROM favorites WHERE user_id = $1 AND post_id = $2 RETURNING user_id, post_id;`

	var userId, postId string
	err := r.db.QueryRow(q, favorite.UserId, favorite.PostId).Scan(&userId, &postId)
	if err != nil {
		log.Printf("Element: %s | Failed to delete favorite in db: %v", element, err)
		return "", "", err
	}

	key := fmt.Sprintf("%s:%s", favorite.UserId, favorite.PostId)
	_, err = r.redis.Del(context.Background(), key).Result()
	if err != nil {
		log.Printf("Element: %s | Failed to delete favorite in redis: %v", element, err)
	}

	return userId, postId, nil
}

func (r *favoriteRepository) Find(startIndex, pageSize int) ([]*models.Favorite, error) {
	q := `SELECT * FROM favorites LIMIT $1 OFFSET $2;`

	var favorites []*models.Favorite
	if err := r.db.Select(&favorites, q, startIndex, pageSize); err != nil {
		log.Printf("Element: %s | Failed to find favorites in db: %v", element, err)
		return nil, err
	}
	if len(favorites) == 0 {
		return []*models.Favorite{}, nil
	}
	return favorites, nil
}

func (r *favoriteRepository) Count(postID string) (int32, error) {
/*	exists, err := r.redis.Get(context.Background(), "*:"+postID).Bool() 
	if exists {
		count, err := r.countRedis(postID)
		if err != nil {
			log.Printf("Element: %s | Failed to count favorites in redis: %v", element, err)
		}
		return count, nil
	}
	*/
	q := `SELECT COUNT(*) FROM favorites WHERE post_id = $1;`

	var count int32
	err := r.db.QueryRow(q, postID).Scan(&count)
	if err != nil {
		log.Printf("Element: %s | Failed to count favorites in db: %v", element, err)
		return 0, err
	}
	return count, nil
}
/*
func (r *favoriteRepository) countRedis(postId string) (int32, error) {
	var totalCount int64
	var cursor uint64

	for {
		keys, _, err := r.redis.Scan(context.Background(), cursor, "*:"+postId, 0).Result()
		if err != nil {
			return 0, err
		}

		if len(keys) == 0 {
			break
		}

		for _, key := range keys {
			val, err := r.redis.Get(context.Background(), key).Int64()
			if err != nil {
				return 0, err
			}
			totalCount += val
		}

		cursor = uint64(len(keys))
	}
	return int32(totalCount), nil
}
*/
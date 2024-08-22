package repositories

// redis: key structure -> user_id:post_id

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/dusk-chancellor/mego-favorite/internal/models"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
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
	redis *redis.Client
}

func NewFavoriteRepository(db *sqlx.DB, rdb *redis.Client) FavoriteRepository {
	return &favoriteRepository{
		db: db,
		redis: rdb,
	}
}

func (r *favoriteRepository) Add(favorite models.Favorite) (string, string, error) {
	q := `INSERT INTO favorites (user_id, post_id) VALUES ($1, $2) RETURNING user_id, post_id;`

	var userId, postId string
	err := r.db.QueryRow(q, favorite.UserId, favorite.PostID).Scan(&userId, &postId)
	if err != nil {
		log.Printf("Element: %s | Failed to add favorite in db: %v", element, err)
		return "", "", err
	}

	key := fmt.Sprintf("%s:%s", favorite.UserId, favorite.PostID)
	_, err = r.redis.Set(context.Background(), key, 1, 24*time.Hour).Result()
	if err != nil {
		log.Printf("Element: %s | Failed to add favorite in redis: %v", element, err)
	}

	return userId, postId, nil
}

func (r *favoriteRepository) Exists(favorite models.Favorite) bool {
	key := fmt.Sprintf("%s:%s", favorite.UserId, favorite.PostID)
	exists, _ := r.redis.Get(context.Background(), key).Bool()
	if exists {
		return true
	}

	q := `SELECT EXISTS(SELECT 1 FROM favorites WHERE user_id = $1 AND post_id = $2);`

	err := r.db.QueryRow(q, favorite.UserId, favorite.PostID).Scan(&exists)
	if err != nil {
		log.Printf("Element: %s | Failed to check if favorite exists in db: %v", element, err)
		return false
	}
	return exists
}

func (r *favoriteRepository) Delete(favorite models.Favorite) (string, string, error) {
	q := `DELETE FROM favorites WHERE user_id = $1 AND post_id = $2 RETURNING user_id, post_id;`

	var userId, postId string
	err := r.db.QueryRow(q, favorite.UserId, favorite.PostID).Scan(&userId, &postId)
	if err != nil {
		log.Printf("Element: %s | Failed to delete favorite in db: %v", element, err)
		return "", "", err
	}

	key := fmt.Sprintf("%s:%s", favorite.UserId, favorite.PostID)
	_, err = r.redis.Del(context.Background(), key).Result()
	if err != nil {
		log.Printf("Element: %s | Failed to delete favorite in redis: %v", element, err)
	}

	return userId, postId, nil
}

func (r *favoriteRepository) Count(postID string) int32 {
/*	exists, err := r.redis.Get(context.Background(), "*:"+postID).Bool() 
	if exists {
		count, err := r.CountRedis(postID)
		if err != nil {
			log.Printf("Element: %s | Failed to count favorites in redis: %v", element, err)
		}
		return count
	}
	*/
	q := `SELECT COUNT(*) FROM favorites WHERE post_id = $1;`

	var count int32
	err := r.db.QueryRow(q, postID).Scan(&count)
	if err != nil {
		log.Printf("Element: %s | Failed to count favorites in db: %v", element, err)
		return 0
	}
	return count
}

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

package models

type Favorite struct {
	PostID string `db:"post_id"`
	Count  int32  `db:"count"`
}

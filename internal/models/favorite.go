package models

type Favorite struct {
	Id	   string `db:"id"`
	UserId string `db:"user_id"`
	PostID string `db:"post_id"`
}

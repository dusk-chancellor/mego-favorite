package models

type Favorite struct {
	Id	   int 	  `db:"id"`
	UserId string `db:"user_id"`
	PostID string `db:"post_id"`
}

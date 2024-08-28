package dto

import (
	"github.com/antibomberman/mego-protos/gen/go/favorite"
	"github.com/dusk-chancellor/mego-favorite/internal/models"
)

func ToPbFavorites(favorites []*models.Favorite) (pbFavorites []*favorite.Favorite) {
	for _, fav := range favorites {
		pbFavorite := &favorite.Favorite{
			UserId: fav.UserId,
			PostId: fav.PostId,
		}
		pbFavorites = append(pbFavorites, pbFavorite)
	}
	return
}
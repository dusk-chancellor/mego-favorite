package grpc

import (
	"context"
	"log"

	pb "github.com/antibomberman/mego-protos/gen/go/favorite"
	"github.com/dusk-chancellor/mego-favorite/internal/models"
	"google.golang.org/grpc/status"
)

const (
	element = "favorite_handlers"
)

func (s *serverAPI) Exists(ctx context.Context, req *pb.ExistsRequest) (*pb.ExistsResponse, error) {
	userID := req.GetUserId()
	postID := req.GetPostId()
	favorite := models.Favorite{UserId: userID, PostID: postID}
	exists := s.service.Exists(favorite)
	return &pb.ExistsResponse{Exists: exists}, nil
}

func (s *serverAPI) Add(ctx context.Context, req *pb.AddRequest) (*pb.AddResponse, error) {
	getUserID := req.GetUserId()
	getPostID := req.GetPostId()
	favorite := models.Favorite{UserId: getUserID, PostID: getPostID}
	userId, postId, err := s.service.Add(favorite)
	if err != nil {
		log.Printf("Element: %s | Failed to add favorite: %v", element, err)
		return nil, status.Error(status.Code(err), "Failed to add favorite")
	}
	return &pb.AddResponse{UserId: userId, PostId: postId}, nil
}

func (s *serverAPI) Delete(ctx context.Context, req *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	getUserID := req.GetUserId()
	getPostID := req.GetPostId()
	favorite := models.Favorite{UserId: getUserID, PostID: getPostID}
	userId, postId, err := s.service.Delete(favorite)
	if err != nil {
		log.Printf("Element: %s | Failed to delete favorite: %v", element, err)
		return nil, status.Error(status.Code(err), "Failed to delete favorite")
	}
	return &pb.DeleteResponse{UserId: userId, PostId: postId}, nil
}

func (s *serverAPI) Count(ctx context.Context, req *pb.CountRequest) (*pb.CountResponse, error) {
	postID := req.GetPostId()
	count := s.service.Count(postID)
	return &pb.CountResponse{Count: count}, nil
}
/*
func (s *serverAPI) Find(ctx context.Context, req *pb.FindRequest) (*pb.FindResponse, error) {

}
*/
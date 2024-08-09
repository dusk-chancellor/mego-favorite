package grpc

import (
	"context"
	"log"

	pb "github.com/antibomberman/mego-protos/gen/go/favorite"
	"google.golang.org/grpc/status"
)

const (
	element = "favorite_handlers"
)

func (s *serverAPI) Exists(ctx context.Context, req *pb.ExistsRequest) (*pb.ExistsResponse, error) {
	postID := req.GetPostId()
	exists := s.service.Exists(postID)
	return &pb.ExistsResponse{Exists: exists}, nil
}

func (s *serverAPI) Add(ctx context.Context, req *pb.AddRequest) (*pb.AddResponse, error) {
	postID := req.GetPostId()
	id, err := s.service.Add(postID)
	if err != nil {
		log.Printf("Element: %s | Failed to add favorite: %v", element, err)
		return nil, status.Error(status.Code(err), "Failed to add favorite")
	}
	return &pb.AddResponse{PostId: id}, nil
}

func (s *serverAPI) Delete(ctx context.Context, req *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	postID := req.GetPostId()
	id, err := s.service.Delete(postID)
	if err != nil {
		log.Printf("Element: %s | Failed to delete favorite: %v", element, err)
		return nil, status.Error(status.Code(err), "Failed to delete favorite")
	}
	return &pb.DeleteResponse{PostId: id}, nil
}

func (s *serverAPI) Count(ctx context.Context, req *pb.CountRequest) (*pb.CountResponse, error) {
	postID := req.GetPostId()
	count := s.service.Count(postID)
	return &pb.CountResponse{Count: count}, nil
}

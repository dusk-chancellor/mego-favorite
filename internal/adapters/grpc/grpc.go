package grpc

import (
	pb "github.com/antibomberman/mego-protos/gen/go/favorite"
	"github.com/dusk-chancellor/mego-favorite/internal/config"
	"github.com/dusk-chancellor/mego-favorite/internal/services"
	"google.golang.org/grpc"
)

type serverAPI struct {
	pb.UnimplementedFavoriteServiceServer

	cfg     *config.Config
	service services.FavoriteService
}

func Register(gRPC *grpc.Server, cfg *config.Config, service services.FavoriteService) {
	pb.RegisterFavoriteServiceServer(gRPC, &serverAPI{
		cfg:     cfg,
		service: service,
	})
}

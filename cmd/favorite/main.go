package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	adapter "github.com/dusk-chancellor/mego-favorite/internal/adapters/grpc"
	"github.com/dusk-chancellor/mego-favorite/internal/config"
	"github.com/dusk-chancellor/mego-favorite/internal/database"
	"github.com/dusk-chancellor/mego-favorite/internal/repositories"
	"github.com/dusk-chancellor/mego-favorite/internal/services"
	"google.golang.org/grpc"
)


func main() {
	cfg := config.LoadConfig()
	db, err := database.ConnectToDB(cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	favoriteRepo := repositories.NewFavoriteRepository(db)
	favoriteLocalCache := services.NewFavoriteLocalCache()
	favoriteService := services.NewFavoriteService(favoriteRepo, favoriteLocalCache)

	l, err := net.Listen("tcp", ":"+cfg.GRPCPort)
	if err != nil {
		log.Fatal(err)
	}

	gRPCServer := grpc.NewServer()
	adapter.Register(gRPCServer, cfg, favoriteService)
	go log.Fatal(gRPCServer.Serve(l))
	
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop
	gRPCServer.GracefulStop()
}

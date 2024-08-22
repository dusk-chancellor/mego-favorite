package main

import (
	"context"
	"fmt"
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
	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc"
)


func main() {
	cfg := config.LoadConfig()
	db, err := database.ConnectToDB(cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.RedisHost, cfg.RedisPort),
		Password: "",
		DB:       0,
	})

	err = rdb.Ping(context.Background()).Err()
	if err != nil {
		log.Fatal("redis connection error")
	}
	log.Println("Connected to Redis")

	favoriteRepo := repositories.NewFavoriteRepository(db, rdb)
	favoriteService := services.NewFavoriteService(favoriteRepo)

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

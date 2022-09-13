package main

import (
	"github.com/NeverlandMJ/GRUD/grud-service/config"
	"github.com/NeverlandMJ/GRUD/grud-service/grudpb"
	"github.com/NeverlandMJ/GRUD/grud-service/repository/database"
	"github.com/NeverlandMJ/GRUD/grud-service/server"
	"github.com/NeverlandMJ/GRUD/grud-service/service"
	"google.golang.org/grpc"
	"net"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}
	repo, err := database.NewPostgresDB(cfg)
	if err != nil {
		panic(err)
	}
	svc := service.NewService(repo)
	srv := server.New(svc)

	lis, err := net.Listen("tcp", net.JoinHostPort(cfg.Host, cfg.Port))
	if err != nil {
		panic(err)
	}
	grpcServer := grpc.NewServer()

	grudpb.RegisterGrudServiceServer(grpcServer, srv)
	if err = grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}

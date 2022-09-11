package main

import (
	"net"

	"github.com/NeverlandMJ/GRUD/collecter-service/collectpb"
	"github.com/NeverlandMJ/GRUD/collecter-service/config"
	"github.com/NeverlandMJ/GRUD/collecter-service/repository/database"
	"github.com/NeverlandMJ/GRUD/collecter-service/server"
	"github.com/NeverlandMJ/GRUD/collecter-service/service"
	"google.golang.org/grpc"
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
	server := server.New(svc)

	lis, err := net.Listen("tcp", net.JoinHostPort(cfg.Host, cfg.Port))
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()
	collectpb.RegisterCollecterServiceServer(grpcServer, server)
	if err = grpcServer.Serve(lis); err != nil {
		panic(err)
	}

}

package server

import "github.com/NeverlandMJ/GRUD/grud-service/service"

type Server struct {
	// grudpb.UnimplementedGrudServiceServer
	svc service.Service
}

func New(svc service.Service) Server {
	return Server{
		svc: svc,
	}
}

package service

import db "github.com/NeverlandMJ/GRUD/grud-service/repository/database"

type Service struct {
	repo db.Repository
}

func NewService(repo db.Repository) Service {
	return Service{
		repo: repo,
	}
}


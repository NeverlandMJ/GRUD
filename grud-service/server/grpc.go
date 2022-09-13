package server

import (
	"context"
	"errors"
	"github.com/NeverlandMJ/GRUD/grud-service/errs"
	"github.com/NeverlandMJ/GRUD/grud-service/grudpb"
	"github.com/NeverlandMJ/GRUD/grud-service/service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	grudpb.UnimplementedGrudServiceServer
	svc service.Service
}

func New(svc service.Service) Server {
	return Server{
		svc: svc,
	}
}

func (s Server) GetPostsByUserID(ctx context.Context, req *grudpb.GetUserPostsRequest) (*grudpb.GetUserPostsResponse, error) {
	userID := int(req.GetUserId())

	posts, err := s.svc.GetPostsByUserID(ctx, userID)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	grudPosts := make([]*grudpb.Data, 0, len(posts))
	for _, p := range posts {
		grudPosts = append(grudPosts, &grudpb.Data{
			Id:     int32(p.ID),
			UserId: int32(p.UserID),
			Title:  p.Title,
			Body:   p.Body,
		})
	}

	return &grudpb.GetUserPostsResponse{
		Datas: grudPosts,
	}, nil
}

func (s Server) GetPostByID(ctx context.Context, req *grudpb.GetPostByIDRequest) (*grudpb.Data, error) {
	postID := int(req.GetPostId())

	post, err := s.svc.GetPostByID(ctx, postID)
	if err != nil {
		if errors.Is(err, errs.ErrPostNotFound) {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, err
	}

	return &grudpb.Data{
		Id:     int32(post.ID),
		UserId: int32(post.UserID),
		Title:  post.Title,
		Body:   post.Body,
	}, nil
}
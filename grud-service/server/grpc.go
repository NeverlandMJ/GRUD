package server

import (
	"context"
	"errors"
	"github.com/NeverlandMJ/GRUD/grud-service/errs"
	"github.com/NeverlandMJ/GRUD/grud-service/grudpb"
	"github.com/NeverlandMJ/GRUD/grud-service/service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
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

func (s Server) ListPosts(ctx context.Context, req *grudpb.ListPostsRequest) (*grudpb.ListPostsResponse, error) {
	posts, all, err := s.svc.ListPosts(ctx, int(req.Page), int(req.Limit))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	grudPosts := make([]*grudpb.Post, 0, len(posts))
	for _, p := range posts {
		grudPosts = append(grudPosts, &grudpb.Post{
			Id:     int32(p.ID),
			UserId: int32(p.UserID),
			Title:  p.Title,
			Body:   p.Body,
		})
	}

	return &grudpb.ListPostsResponse{
		Data: &grudpb.GetPostsResponse{
			Posts: grudPosts,
		},
		All: int32(all),
	}, nil
}

func (s Server) GetPostsByUserID(ctx context.Context, req *grudpb.GetUserPostsRequest) (*grudpb.GetPostsResponse, error) {
	userID := int(req.GetUserID())

	posts, err := s.svc.GetPostsByUserID(ctx, userID)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	grudPosts := make([]*grudpb.Post, 0, len(posts))
	for _, p := range posts {
		grudPosts = append(grudPosts, &grudpb.Post{
			Id:     int32(p.ID),
			UserId: int32(p.UserID),
			Title:  p.Title,
			Body:   p.Body,
		})
	}

	return &grudpb.GetPostsResponse{
		Posts: grudPosts,
	}, nil
}

func (s Server) GetPostByID(ctx context.Context, req *grudpb.GetPostByIDRequest) (*grudpb.Post, error) {
	postID := int(req.GetPostID())

	post, err := s.svc.GetPostByID(ctx, postID)
	if err != nil {
		if errors.Is(err, errs.ErrPostNotFound) {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &grudpb.Post{
		Id:     int32(post.ID),
		UserId: int32(post.UserID),
		Title:  post.Title,
		Body:   post.Body,
	}, nil
}

func (s Server) DeletePost(ctx context.Context, req *grudpb.DeletePostRequest) (*emptypb.Empty, error) {
	postID := int(req.GetPostID())

	err := s.svc.DeletePost(ctx, postID)
	if err != nil {
		return &emptypb.Empty{}, status.Error(codes.Internal, err.Error())
	}

	return &emptypb.Empty{}, nil
}

func (s Server) UpdateTitle(ctx context.Context, req *grudpb.UpdateTitleRequest) (*grudpb.Post, error) {
	post, err := s.svc.UpdateTitle(ctx, int(req.PostID), req.NewTitle)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &grudpb.Post{
		Id:     int32(post.ID),
		UserId: int32(post.UserID),
		Title:  post.Title,
		Body:   post.Body,
	}, nil
}
func (s Server) UpdateBody(ctx context.Context, req *grudpb.UpdateBodyRequest) (*grudpb.Post, error) {
	post, err := s.svc.UpdateBody(ctx, int(req.PostID), req.NewBody)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &grudpb.Post{
		Id:     int32(post.ID),
		UserId: int32(post.UserID),
		Title:  post.Title,
		Body:   post.Body,
	}, nil
}

package grpc_server

import (
	"context"
	"github.com/lifedaemon-kill/ozon-url-shortener-api/internal/service"
	pb "github.com/lifedaemon-kill/ozon-url-shortener-api/pkg/grpc/gen/url_shortener"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log/slog"
)

func NewURLService(repo service.UrlShortener, log *slog.Logger) pb.URLServiceServer {
	return &urlService{service: repo, log: log}
}

type urlService struct {
	pb.UnimplementedURLServiceServer
	service service.UrlShortener
	log     *slog.Logger
}

func Registration(server *grpc.Server, urlService pb.URLServiceServer) {
	pb.RegisterURLServiceServer(server, urlService)
	reflection.Register(server)
}

func (s *urlService) SaveURL(ctx context.Context, req *pb.SaveURLRequest) (*pb.SaveURLResponse, error) {
	alias, err := s.service.CreateAlias(req.Source_URL)
	if err != nil {
		return nil, err
	}
	return &pb.SaveURLResponse{Alias_URL: alias}, nil
}

func (s *urlService) FetchURL(ctx context.Context, req *pb.FetchURLRequest) (*pb.FetchURLResponse, error) {
	source, err := s.service.FetchSource(req.Alias_URL)
	if err != nil {
		return nil, err
	}
	return &pb.FetchURLResponse{Source_URL: source}, nil
}

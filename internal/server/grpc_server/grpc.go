package grpc_server

import (
	"context"
	"github.com/lifedaemon-kill/ozon-url-shortener-api/internal/service"
	pb "github.com/lifedaemon-kill/ozon-url-shortener-api/pkg/grpc/gen/url_shortener"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log/slog"
	"time"
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
	s.log.Debug("grpc.SaveURL", "Request", req)

	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	alias, err := s.service.CreateAlias(ctx, req.Source_URL)
	if err != nil {
		s.log.Error("grpc.SaveURL", "err", err)
		return nil, err
	}

	s.log.Debug("grpc.SaveURL success", "alias", alias)
	return &pb.SaveURLResponse{Alias_URL: alias}, nil
}

func (s *urlService) FetchURL(ctx context.Context, req *pb.FetchURLRequest) (*pb.FetchURLResponse, error) {
	s.log.Debug("grpc.FetchURL", "Request", req)

	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	source, err := s.service.FetchSource(ctx, req.Alias_URL)
	if err != nil {
		s.log.Error("grpc.FetchURL", "err", err)
		return nil, err
	}
	s.log.Debug("grpc.FetchURL success", "source", source)
	return &pb.FetchURLResponse{Source_URL: source}, nil
}

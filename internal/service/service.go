package service

import (
	"context"
	"errors"
	"github.com/lifedaemon-kill/ozon-url-shortener-api/internal/pkg/config"
	"github.com/lifedaemon-kill/ozon-url-shortener-api/internal/pkg/internal_errors"
	"github.com/lifedaemon-kill/ozon-url-shortener-api/internal/pkg/lib"
	"github.com/lifedaemon-kill/ozon-url-shortener-api/internal/storage"
	"log/slog"
)

type UrlShortener interface {
	CreateAlias(ctx context.Context, sourceURL string) (aliasURL string, err error)
	FetchSource(ctx context.Context, aliasURL string) (sourceURL string, err error)
}

type service struct {
	repo      storage.Storage
	log       *slog.Logger
	genConfig config.URLGenerator
	host      string
}

func New(conf config.URLGenerator, repo storage.Storage, host string, log *slog.Logger) UrlShortener {
	return &service{
		repo:      repo,
		log:       log,
		genConfig: conf,
		host:      host,
	}
}

func (s *service) CreateAlias(ctx context.Context, sourceURL string) (string, error) {
	s.log.Debug("start service.CreateAlias", "sourceURL", sourceURL)
	//Валидация ссылки
	isUrl := lib.IsURL(sourceURL)
	if !isUrl {
		s.log.Error("parsing link in service.CreateAlias was failed", "source", sourceURL)
		return "", ierrors.InvalidURL
	}

	//Генерируем новую ссылку
	alias := lib.GenerateLinkStrBuilder(s.genConfig.URLLength, s.genConfig.AllowedSymbols)

	//проверяем, есть ли уже такая ссылка в базе
	_, err := s.repo.FetchURL(ctx, alias)

	//В случае если произошла коллизия, нужно перегенерировать ссылку
	for err == nil {
		alias = lib.GenerateLinkStrBuilder(s.genConfig.URLLength, s.genConfig.AllowedSymbols)
		_, err = s.repo.FetchURL(ctx, alias)
	}
	if errors.Is(err, ierrors.NoSuchValue) {
		//Пытаемся записать, если данная строка есть, должна вернуться ошибка source alredy exist
		err = s.repo.SaveURL(ctx, sourceURL, alias)

		if errors.Is(err, ierrors.SourceAlreadyExist) {
			s.log.Error("source already exists", "sourceURL", sourceURL)
			return "", ierrors.SourceAlreadyExist
		}

		s.log.Debug("end service.CreateAlias", "sourceURL", sourceURL, "alias", alias)
		return alias, nil
	}
	//Иначе что-то отвалилось в бд...
	s.log.Error("end service.CreateAlias was failed", "err", err)
	return "", err
}

func (s *service) FetchSource(ctx context.Context, aliasURL string) (sourceURL string, err error) {
	s.log.Debug("start service.FetchSource", "alias", aliasURL)

	source, err := s.repo.FetchURL(ctx, aliasURL)
	if err != nil {
		s.log.Error("end service.CreateAlias was failed", "err", err)
		return "", err
	}
	s.log.Debug("end service.FetchSource", "sourceURL", source, "alias", aliasURL)
	return s.host + source, nil
}

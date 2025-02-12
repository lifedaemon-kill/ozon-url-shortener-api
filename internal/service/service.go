package service

import (
	"errors"
	"github.com/lifedaemon-kill/ozon-url-shortener-api/internal/pkg/config"
	"github.com/lifedaemon-kill/ozon-url-shortener-api/internal/pkg/internal_errors"
	"github.com/lifedaemon-kill/ozon-url-shortener-api/internal/pkg/lib"
	"github.com/lifedaemon-kill/ozon-url-shortener-api/internal/storage"
	"log/slog"
)

type UrlShortener interface {
	CreateAlias(sourceURL string) (aliasURL string, err error)
	FetchSource(aliasURL string) (sourceURL string, err error)
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

func (s *service) CreateAlias(sourceURL string) (string, error) {
	//Генерируем новую ссылку
	alias := lib.GenerateLinkStrBuilder(s.genConfig.URLLength, s.genConfig.AllowedSymbols)

	//Проверяем, есть ли такая уже в бд
	_, err := s.repo.FetchURL(alias)

	//В случае если произошла коллизия, нужно перегенерировать ссылку
	for err == nil {
		alias = lib.GenerateLinkStrBuilder(s.genConfig.URLLength, s.genConfig.AllowedSymbols)
		_, err = s.repo.FetchURL(alias)
	}

	//Проверяем по ошибке, что в бд нет такой строки, и добавляем
	if errors.Is(err, ierrors.NoSuchValue) {
		if err = s.repo.SaveURL(sourceURL, alias); err != nil {
			return "", err
		}
		return alias, nil
	}
	//Иначе что-то отвалилось в бд...
	return "", err

}

func (s *service) FetchSource(aliasURL string) (sourceURL string, err error) {
	source, err := s.repo.FetchURL(aliasURL)
	if err != nil {
		return "", err
	}
	return s.host + source, nil
}

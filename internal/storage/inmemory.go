package storage

import "github.com/lifedaemon-kill/ozon-url-shortener-api/internal/pkg/internal_errors"

type inMemoryStorage struct {
	db map[string]string
}

func NewInMemory() Storage {
	db := make(map[string]string)
	return &inMemoryStorage{
		db: db,
	}
}

func (i *inMemoryStorage) SaveURL(sourceURL, aliasURL string) error {
	res := i.db[aliasURL]
	if res == "" {
		return ierrors.ValueAlreadyExist
	}
	return nil
}

func (i *inMemoryStorage) FetchURL(aliasURL string) (sourceURL string, err error) {
	//TODO implement me
	panic("implement me")
}

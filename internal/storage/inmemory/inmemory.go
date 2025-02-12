package inmemory

import (
	"github.com/lifedaemon-kill/ozon-url-shortener-api/internal/pkg/internal_errors"
	"github.com/lifedaemon-kill/ozon-url-shortener-api/internal/storage"
	"sync"
)

type inMemoryStorage struct {
	sourceToAlias map[string]string
	aliasToSource map[string]string
	mu            sync.Mutex
}

func New() storage.Storage {
	return &inMemoryStorage{
		sourceToAlias: make(map[string]string),
		aliasToSource: make(map[string]string),
		mu:            sync.Mutex{},
	}
}

// SaveURL if url already exist, then method should return ierror.ValueAlreadyExist
func (i *inMemoryStorage) SaveURL(sourceURL, aliasURL string) error {
	i.mu.Lock()
	res := i.sourceToAlias[sourceURL]
	if res != "" {
		i.mu.Unlock()
		return ierrors.ValueAlreadyExist
	}
	i.sourceToAlias[sourceURL] = aliasURL
	i.aliasToSource[aliasURL] = sourceURL
	i.mu.Unlock()

	return nil
}

// FetchURL if no such url, then method should return ierror.NoSuchValue
func (i *inMemoryStorage) FetchURL(aliasURL string) (string, error) {
	i.mu.Lock()
	source := i.aliasToSource[aliasURL]
	if source == "" {
		i.mu.Unlock()
		return "", ierrors.NoSuchValue
	}
	i.mu.Unlock()

	return source, nil
}

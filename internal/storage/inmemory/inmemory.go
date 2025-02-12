package inmemory

import (
	"context"
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

// SaveURL если ссылка уже существует, возвращает ошибку ierror.SourceAlreadyExist
func (i *inMemoryStorage) SaveURL(ctx context.Context, sourceURL, aliasURL string) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	i.mu.Lock()
	defer i.mu.Unlock()

	_, exist := i.sourceToAlias[sourceURL]
	if exist {
		return ierrors.SourceAlreadyExist
	}

	i.sourceToAlias[sourceURL] = aliasURL
	i.aliasToSource[aliasURL] = sourceURL

	return nil
}

// FetchURL если не такой ссылки, то возвращает ierror.NoSuchValue
func (i *inMemoryStorage) FetchURL(ctx context.Context, aliasURL string) (string, error) {
	select {
	case <-ctx.Done():
		return "", ctx.Err()
	default:
	}

	i.mu.Lock()
	defer i.mu.Unlock()

	source, exist := i.aliasToSource[aliasURL]
	if !exist {
		return "", ierrors.NoSuchValue
	}

	return source, nil
}

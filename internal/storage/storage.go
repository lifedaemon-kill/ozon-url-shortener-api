package storage

import "context"

type Storage interface {
	// SaveURL if url already exist? then method should return ierror.ValueAlreadyExist
	SaveURL(ctx context.Context, sourceURL, aliasURL string) error

	// FetchURL if no such url, then method should return ierror.NoSuchValue
	FetchURL(ctx context.Context, aliasURL string) (sourceURL string, err error)
}

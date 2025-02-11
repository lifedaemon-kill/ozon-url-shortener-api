package storage

type Storage interface {
	SaveURL(sourceURL, aliasURL string) error

	// FetchURL if no such url, then method must to response internal_errors.NoSuchValue
	FetchURL(aliasURL string) (sourceURL string, err error)
}

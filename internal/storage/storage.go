package storage

type Storage interface {
	// SaveURL if url already exist? then method should return ierror.ValueAlreadyExist
	SaveURL(sourceURL, aliasURL string) error

	// FetchURL if no such url, then method should return ierror.NoSuchValue
	FetchURL(aliasURL string) (sourceURL string, err error)
}

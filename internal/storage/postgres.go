package storage

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/lifedaemon-kill/ozon-url-shortener-api/internal/pkg/config"
)

type postgres struct {
	db *sqlx.DB
}

func NewPostgres(conf config.DB) (Storage, error) {
	db, err := sqlx.Connect("postgres", conf.DSN)
	if err != nil {
		return nil, err
	}
	return &postgres{
		db: db,
	}, nil
}
func (p *postgres) SaveURL(sourceURL, aliasURL string) error {
	//TODO implement me
	panic("implement me")
}

func (p *postgres) FetchURL(aliasURL string) (sourceURL string, err error) {
	//TODO implement me
	panic("implement me")
}

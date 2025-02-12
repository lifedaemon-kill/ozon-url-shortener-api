package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/lifedaemon-kill/ozon-url-shortener-api/internal/pkg/config"
	ierrors "github.com/lifedaemon-kill/ozon-url-shortener-api/internal/pkg/internal_errors"
	"github.com/lifedaemon-kill/ozon-url-shortener-api/internal/storage"
	"time"
)

type postgres struct {
	db *sqlx.DB
}

func New(conf config.DB) (storage.Storage, error) {
	db, err := sqlx.Connect("postgres", fmt.Sprintf(conf.DSN))
	if err != nil {
		return nil, err
	}
	return &postgres{
		db: db,
	}, nil
}
func (p *postgres) SaveURL(sourceURL, aliasURL string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	query := `INSERT INTO urls (sourceURL, aliasURL) VALUES ($1, $2)`

	_, err := p.db.ExecContext(ctx, query, sourceURL, aliasURL)
	if err != nil {
		return err
	}
	return nil
}

func (p *postgres) FetchURL(aliasURL string) (sourceURL string, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	query := `SELECT sourceURL FROM urls WHERE aliasURL=$1`

	row := p.db.QueryRowContext(ctx, query, aliasURL)
	err = row.Scan(&sourceURL)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", ierrors.NoSuchValue
		}
		return "", err
	}
	return sourceURL, nil
}

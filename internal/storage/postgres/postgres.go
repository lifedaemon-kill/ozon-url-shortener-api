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

func (p *postgres) SaveURL(ctx context.Context, sourceURL, aliasURL string) error {
	query := `
		INSERT INTO urls (sourceURL, aliasURL) 
		VALUES ($1, $2)
		ON CONFLICT (sourceURL)
		`

	result, err := p.db.ExecContext(ctx, query, sourceURL, aliasURL)
	if err != nil {
		return err
	}

	// Проверяем, была ли вставка успешной
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	// Если rowsAffected == 0, значит запись уже существует
	if rowsAffected == 0 {
		return ierrors.SourceAlreadyExist
	}

	return nil
}

func (p *postgres) FetchURL(ctx context.Context, aliasURL string) (sourceURL string, err error) {
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

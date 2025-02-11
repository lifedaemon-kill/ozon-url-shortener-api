-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS urls
(
    id         SERIAL PRIMARY KEY,
    source_url TEXT NOT NULL UNIQUE,
    alias_url  TEXT NOT NULL
);

CREATE INDEX indx_urls_alias_url on urls (alias_url);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS urls;
DROP INDEX IF EXISTS indx_urls_alias_url;
-- +goose StatementEnd

package lib

import (
	"github.com/lifedaemon-kill/ozon-url-shortener-api/internal/pkg/internal_errors"
)

const (
	Postgres = "postgres"
	InMemory = "inmemory"
)

// ParseStorageType ищет тэг -s и следующее за ним слово,
// Если оно соответствует константам, возвращает его, иначе ошибка: ierrors.UnknownStorageType
//
// Константы: Postgres / InMemory
//
// В случае если тэга -s нет, по умолчанию возвращает Postgres
func ParseStorageType(args []string) (storageType string, err error) {
	if len(args) < 2 {
		return Postgres, nil
	}
	for i := range args[:len(args)-1] {
		if args[i] == "-s" {
			switch args[i+1] {
			case Postgres:
				return Postgres, nil
			case InMemory:
				return InMemory, nil
			default:
				return "", ierrors.UnknownStorageType
			}
		}
	}
	return Postgres, nil
}

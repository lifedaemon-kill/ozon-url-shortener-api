package lib

import (
	"errors"
	"fmt"
)

// ParseStorageType парсит аргументы и ищет тег -s
// и возвращает следующее слово, если оно соответствует константе
// ЕСЛИ ТЕГ НЕ НАЙДЕН, ВОЗВРАЩАЕТ POSTGRES
const (
	postgres = "postgres"
	memory   = "memory"
)

func ParseStorageType(args []string) (storageType string, err error) {
	for i := range args[:len(args)-1] {
		fmt.Println(i, args[i])
		if args[i] == "-s" {
			switch args[i+1] {
			case postgres:
				return postgres, nil
			case memory:
				return memory, nil
			default:
				return "", errors.New("unknown storage type")
			}
		}
	}
	return postgres, nil
}

package lib

import (
	"errors"
	"fmt"
	"github.com/go-playground/assert/v2"
	ierrors "github.com/lifedaemon-kill/ozon-url-shortener-api/internal/pkg/internal_errors"
	"testing"
)

func TestParseStorageTypeCorrect(t *testing.T) {
	args := [][]string{
		{"-s", "postgres"},
		{"-d", "h"},
		{"-d", "hello", "-s", "postgres"},
		{"-s", "inmemory", "-s", "inmemory", "asdda"},
		{"ddd", "hello", "-s", "inmemory", "asdads"},
	}
	expected := []string{
		Postgres,
		Postgres,
		Postgres,
		InMemory,
		InMemory,
	}

	for i, arg := range args {
		result, err := ParseStorageType(arg)
		if err != nil {
			t.Error(err)
		} else {
			assert.Equal(t, expected[i], result)
		}
		fmt.Println(expected[i], result)
	}
}

func TestParseStorageTypeNotCorrect(t *testing.T) {
	args := [][]string{
		{"-s", "postgresql"},
		{"-s", "-s", "inmemory", "asdda"},
		{"-s", "asdads", "inmemory"},
	}

	for i, arg := range args {
		st, err := ParseStorageType(arg)
		if err != nil {
			assert.Equal(t, ierrors.UnknownStorageType, err)
		} else {
			t.Error(errors.New("вернулось не то что ожидалось"), i, st, err)
		}
		fmt.Println(st)
	}
}

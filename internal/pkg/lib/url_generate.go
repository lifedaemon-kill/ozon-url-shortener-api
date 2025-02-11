package lib

import (
	"math/rand"
	"strings"
	"sync"
)

// GenerateLinkStrBuilder Создает строку через str builder и math rand
// length - длина строки
// symbols - пул допустимых символов
func GenerateLinkStrBuilder(length int, symbols []rune) string {
	var builder strings.Builder
	builder.Grow(length)

	symbolsLen := len(symbols)

	for range length {
		builder.WriteRune(symbols[rand.Intn(symbolsLen)])
	}

	return builder.String()
}

//Another algorithms are deprecated
//-------------------------------------------------------------------

// GenerateLinkLinear линейно
// генерирует случайную строку длины length из пула символов symbols
// Deprecated
func generateLinkLinear(length int, symbols []rune) string {
	result := make([]rune, length)

	for i := range length {
		result[i] = symbols[rand.Intn(len(symbols))]
	}

	return string(result)
}

// generateLinkParallel создает руны в горутинах
// генерирует случайную строку длины length из пула символов symbols
//
// Deprecated
func generateLinkParallel(length int, symbols []rune) string {
	result := make([]rune, length)

	wg := sync.WaitGroup{}
	for i := range length {
		wg.Add(1)
		go func() {
			defer wg.Done()
			result[i] = symbols[rand.Intn(len(symbols))]
		}()
	}
	wg.Wait()
	return string(result)
}

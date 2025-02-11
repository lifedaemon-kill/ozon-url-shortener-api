package lib

import (
	"fmt"
	"github.com/go-playground/assert/v2"
	testifyAssert "github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestIsItWorks(t *testing.T) {
	length := 10
	symbols := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_")

	for range 150 {
		link := GenerateLinkStrBuilder(length, symbols)
		assert.Equal(t, length, len(link))
		for _, char := range link {
			testifyAssert.Contains(t, symbols, char, "Символ %c не содержится в symbols", char)
		}
	}
}

func TestSpeed(t *testing.T) {
	length := 10
	symbols := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_")
	var start, end time.Time

	for range 30 {
		//Linear rune concat
		start = time.Now()
		for i := 0; i < 600_000; i++ {
			arr := generateLinkLinear(length, symbols)
			_ = arr
		}
		end = time.Now()
		linearTime := end.Sub(start).Nanoseconds() / 1_000_000

		//Linear str builder
		start = time.Now()
		for i := 0; i < 600_000; i++ {
			arr := GenerateLinkStrBuilder(length, symbols)
			_ = arr
		}
		end = time.Now()
		strBuilderTime := end.Sub(start).Nanoseconds() / 1_000_000

		/*
			//Parallel rune concat
			start = time.Now()
			for i := 0; i < 600_000; i++ {
				arr := generateLinkParallel(length, symbols)
				_ = arr
			}
			end = time.Now()
			parallelRuneTime := end.Sub(start).Nanoseconds() / 1_000_000
		*/
		fmt.Println(linearTime, strBuilderTime)
	}
	//средние показатели в миллисекундах
	//296 179 6854
	//282 198 6383
	//351 178 6241 параллельность проиграла без шансов далее без нее
	/*
		242 171
		237 159
		228 153
		226 152
		218 166
		233 152
		230 169
		229 151
		225 189
		214 170
		223 159
		230 174
	*/
	//хитрыми вычислениями питона и пандаса получаем
	//Среднее: 227.92, 163.75, 6492.66
	//Три метода помечаю как устаревшие, оставляем только GenerateLinkStrBuilder
}

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
		strBuilder := end.Sub(start).Nanoseconds() / 1_000_000

		start = time.Now()
		for i := 0; i < 600_000; i++ {
			arr := generateLinkStrBuilderWOLEN(length, symbols)
			_ = arr
		}
		end = time.Now()
		build2 := end.Sub(start).Nanoseconds() / 1_000_000
		/*
			//Parallel rune concat
			start = time.Now()
			for i := 0; i < 600_000; i++ {
				arr := generateLinkParallel(length, symbols)
				_ = arr
			}
			end = time.Now()
			parallelRune := end.Sub(start).Nanoseconds() / 1_000_000
		*/
		fmt.Println(linearTime, strBuilder, build2)
	}
	//средние показатели в миллисекундах
	//296 179 172 6854
	//282 198 186 6383
	//351 178 160 6241 параллельность проиграла без шансов далее без нее
	/*
	   373 441 467
	   395 219 212
	   325 194 178
	   266 195 182
	   255 179 186
	   254 180 203
	   253 195 267
	   313 197 185
	   245 198 186
	   252 192 187
	   301 218 178
	   277 202 243
	   274 189 231
	   358 223 223
	   272 203 195
	   254 166 178
	   279 226 264
	   279 159 183
	   277 198 173
	   254 250 198
	   252 166 169
	   303 179 190
	   256 195 185
	   251 200 189
	   254 185 177
	   242 182 175
	   272 173 162
	   239 171 171
	   242 183 166
	   255 203 164
	*/
	//хитрыми вычислениями питона и пандаса получаем
	//Среднее: 277.40, 202.03, 202.23
	//Три метода помечаю как устаревшие, оставляем только GenerateLinkStrBuilder
}

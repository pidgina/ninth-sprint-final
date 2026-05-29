package main

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {
	if size <= 0 {
		log.Println("Передан размер слайса меньше или равный нулю. Дальнейшее выполнение программы некорректно.")
		return nil
	}
	sliceInteger := make([]int, size)

	for i := 0; i < size; i++ {
		rnd := rand.Intn(size)
		sliceInteger[i] = rnd + 1 //  задание требует целые положительные числа

	}

	return sliceInteger
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	if data == nil {
		log.Println("Попытка вычесть максимум из nil слайса. Дальнейшее выполнение программы некорректно.")
		return 0
	}
	if len(data) == 0 {
		log.Println("Попытка вычесть максимум из пустого слайса. Дальнейшее выполнение программы некорректно.")
		return 0
	}

	maxNumberOfRandom := data[0] // так корректнее, мне кажется.

	for _, v := range data {
		if maxNumberOfRandom < v {
			maxNumberOfRandom = v
		}
	}

	return maxNumberOfRandom

}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {

	var wg sync.WaitGroup
	var mu sync.Mutex

	sliceMaxSlice := []int{}        // слайс для хранения максимальных чисел для дальнейшего перебора
	lenSlice := len(data)           // длина слайса
	valueSlice := lenSlice / CHUNKS // размер слайса

	for i := 0; i < 8; i++ {

		startSlice := i * valueSlice        // начальный индекс
		endSlice := startSlice + valueSlice // конечный индекс

		wg.Add(1)
		go func(start, end int) {
			defer wg.Done()

			max := data[0]

			for _, v := range data[start:end] {
				if max < v {
					max = v
				}
			}
			mu.Lock() // блокируем критическую секцию кода
			sliceMaxSlice = append(sliceMaxSlice, max)
			mu.Unlock() // разблокируем критическую секцию кода
		}(startSlice, endSlice)

	}
	wg.Wait()
	return maximum(sliceMaxSlice)
}

func main() {
	// в задаче сказано "Время нахождения максимумов измеряйте в Микросекундах.",
	// однако в шаблоне на github в выводе указаны "ms", изменил на "us".
	// Также в шаблоне некорректно используются "\n" (в printf нет переноса), поправил для лучшего вывода.
	// https://github.com/Yandex-Practicum/go1fl-sprint9-final-tpl/blob/main/main.go - если нужно для проверки.

	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	sliceRand := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	timeStart := time.Now()
	max := maximum(sliceRand)
	times := time.Since((timeStart))
	elapsed := times.Microseconds()
	// можно обойтись без переменной заменив elapsed на "times.Microseconds()" в принте,
	//  но т.к. в задаче стоит elapsed решил просто добавить переменную.
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d us\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	timeStart = time.Now()
	max = maxChunks(sliceRand)
	times = time.Since(timeStart)
	elapsed = times.Microseconds()
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d us\n", max, elapsed)
}

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
    fmt.Println("\n Задача 9")
    // Создаем каналы для передачи данных между этапами конвейера
	numbers := []int{1, 2, 3, 4, 5} 
	inputCh := make(chan int)
	outputCh := make(chan int)

	// Горутина для чтения чисел из массива и отправки их в первый канал
	go func() {
		for _, num := range numbers {
			inputCh <- num
		}
		close(inputCh)
	}()

	// Горутина для умножения чисел на 2 и отправки результата во второй канал
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		for num := range inputCh {
            time.Sleep(500 * time.Millisecond)
			result := num * 2
			outputCh <- result
		}
		close(outputCh)
		wg.Done()
	}()

	// Вывод результатов из второго канала в stdout
    wg.Add(1)
	go func() {
		for result := range outputCh {
			fmt.Println(result)
		}
        wg.Done()
	}()
	wg.Wait()
}

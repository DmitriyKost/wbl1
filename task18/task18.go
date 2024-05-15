package main

import (
	"fmt"
	"sync"
	"time"
)

type Counter struct {
    val int
    mx sync.Mutex
}

// Инкрементирует значение val
//
// Чтобы это было возможно в конкурентной среде, воспользуемся мутексом.
func (c *Counter) Increment() {
    c.mx.Lock()
    c.val++
    time.Sleep(200 * time.Millisecond) // Для более читаемого вывода
    c.mx.Unlock()
}

func (c *Counter) GetVal() int {
    return c.val
}



func main() {
    fmt.Println("\n Задача 18")

    c := Counter{val:0}

    var wg sync.WaitGroup
    // Запустим 5 воркеров, значение счетчика должно стать 50
    for i := 0; i < 5; i++ {
        wg.Add(1)
        go worker(&wg, &c, i+1)
    }
    wg.Wait()

    fmt.Printf("Значение счетчика: %d\n", c.GetVal())
}

// Воркер пытается инкрементировать значение счетчика 10 раз
func worker(wg *sync.WaitGroup, c *Counter, workerID int) {
    defer wg.Done()
    for i := 0; i < 10; i++ {
        fmt.Printf("Воркер %d пытается инкрементировать...\n", workerID)
        c.Increment()
        fmt.Printf("Воркер %d инкрементировал...\n", workerID)
    }
}

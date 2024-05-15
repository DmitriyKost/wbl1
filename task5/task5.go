package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
    fmt.Println("\n Задача 5")
    fmt.Println("Введите N - число секунд.")
    var n time.Duration
    var wg sync.WaitGroup
    fmt.Scanf("%d", &n)

    ch := make(chan interface{})
    wg.Add(2)
    go sender(ch, time.Second * time.Duration(n), &wg)
    go receiver(ch, &wg)
    wg.Wait()
    fmt.Println("Done")
}

// Создаем отправителя, который будет последовательно отправлять данные в канал, в течение указанного периода времени.
func sender(ch chan<- interface{}, duration time.Duration, wg *sync.WaitGroup) {
    defer close(ch)
    defer wg.Done()
    timer := time.NewTimer(duration)
    for {
        select {
        case <-timer.C:
            return
        case ch <- "Some message...":
        }
    }
}

// Получатель будет обрабатывать сообщения до тех пор пока канал не закроется.
func receiver(ch <-chan interface{}, wg *sync.WaitGroup) {
    defer wg.Done()
    for val := range ch {
        fmt.Println("Received: ", val)
        time.Sleep(500 * time.Millisecond)
    }
}

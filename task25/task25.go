package main

import (
	"fmt"
	"time"
)

func sleep(d time.Duration) {
    // After будет ждать указанное время, чтобы послать в канал текущее время,
    // блокируя горутину.
    <-time.After(d)
}

func main() {
    fmt.Println("\n Задача 25")
    fmt.Println("Засыпаем на 3 секунды...")
    sleep(time.Second * 3)
    fmt.Println("Проснулись!")
}

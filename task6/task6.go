package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
    fmt.Println("\n Задача 6")
    firstWay()
    secondWay()
    thirdWay()
}

// В данном способе мы приостанавливаем выполнение главной горутины до тех пор пока не выполнится воркер,
// с помощью небуферизованного канала.
func firstWay() {
    fmt.Println("   Первый способ: канал")
    done := make(chan bool)
    fmt.Println("Ожидаем воркера")
    go workerWithChan(done)
    <- done
    fmt.Println("Ожидание завершено")
}

func workerWithChan(done chan bool) {
    fmt.Println("Воркер начал работу")
    time.Sleep(2 * time.Second)
    fmt.Println("Воркер закончил работу")
    done <- true
}

// Остановка выполнения/прерывание горутины с помощью контекста.
func secondWay() {
    fmt.Println("   Второй способ: контекст")
    // Создаем контекст с таймаутом, чтобы прервать воркера после 2-х секунд.
    ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
    defer cancel()

    go longRunningWorker(ctx)
    go waitingWorker(ctx)

    time.Sleep(4*time.Second)
}

// Воркер, работающий слишком долго, прерываемый контекстом.
func longRunningWorker(ctx context.Context) {
    fmt.Println("Долгий Воркер начал работу")
    // Выполнится та ветка select, которая сработает раньше, т.е. прерывание.
    select {
    case <-time.After(3 * time.Second):
        fmt.Println("Долгий Воркер завершил работу")
    case <-ctx.Done():
        fmt.Println("Долгий Воркер был приостановлен/прерван...")
    }
}

// Воркер начинающий работать только после выполнения контекста.
func waitingWorker(ctx context.Context) {
    fmt.Println("Ждущий воркер ждет")
    <-ctx.Done()
    fmt.Println("Ждущий воркер дождался")
}

// С помощью мутекса
var mx sync.Mutex

func thirdWay() {
    fmt.Println("   Третий способ: мутекс")
    var wg sync.WaitGroup
    wg.Add(2)
    go countingWorker(&wg)
    time.Sleep(1 * time.Second)
    go workerWithMutex(&wg)
    wg.Wait()
}

// Воркер который будет блокировать остальные горутины, чтобы выполнить свою работу
// с помощью мутекса.
func workerWithMutex(wg *sync.WaitGroup) {
    defer wg.Done()
	fmt.Println("Прерывающий воркер пытается прервать с помощью мутекса...")
    mx.Lock()
	fmt.Println("Прерывающий воркер получил доступ к мьютексу и прерывает считающего воркера...")
    time.Sleep(3 * time.Second)
    mx.Unlock()
}

// Воркер который считает секунды (10).
func countingWorker(wg *sync.WaitGroup) {
    defer wg.Done()
    fmt.Println("Считающий воркер считает:")
    for i := 0; i < 10; i++ {
        mx.Lock()
        fmt.Println(i+1)
        time.Sleep(time.Second)
        mx.Unlock()
    }
}

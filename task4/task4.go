package main

import (
    "fmt"
    "os"
    "os/signal"
    "sync"
    "syscall"
    "time"
)

func main() {
    fmt.Println("\n Задача 4")
    fmt.Println("Введите N - число воркеров.")
    var n int
    fmt.Scanf("%d", &n)

    // Создаем канал для ожидания сигнала о прекращении работы
    interruptChannel := make(chan os.Signal, 1)
    signal.Notify(interruptChannel, syscall.SIGINT, syscall.SIGTERM)

    mainChan := make(chan interface{})

    var wg sync.WaitGroup
    for i := 0; i < n; i++ {
        wg.Add(1)
        go worker(&wg, mainChan, i)
    }

    go func(){
        for {
            // Выполняем одно из действий
            select {
            // Eсли получаем ^C то закрываем канал и выходим из цикла.
            // При закрытии канала, воркеры выйдут из цикла чтения for, получив нулевое значение.
            case <-interruptChannel:
                close(mainChan)
                return
            // В любом другом случае отправляем сообщение в канал.
            case mainChan <- "Some message...":
            }
            // Останавливаем горутину, чтобы вывод был читаемым
        }
    }()
    // Дожидаемся выполнения всех воркеров.
    wg.Wait()
    // Таким образом основная горутина блокируется до тех пор пока все воркеры не выполнят свою работу,
    // реализуется "graceful shutdown", т.е. воркеры успеют выполнить "незаконченные дела" перед остановкой приложения.
    fmt.Println("Received ^C. Shutting down...")
}

func worker(wg *sync.WaitGroup, ch chan interface{}, workerID int) {
    defer wg.Done()
    for data := range ch {
        fmt.Printf("Worker %d received: %v\n", workerID+1, data)
        // Останавливаем горутину, чтобы вывод был читаемым
        time.Sleep(1000 * time.Millisecond)
    }
    // Когда закроется канал воркер выйдет из цикла и сообщит об этом
    fmt.Printf("Worker %d exiting...\n", workerID)
}

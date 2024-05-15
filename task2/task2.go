package main

import (
    "fmt"
    "sync"
)

func main() {
    fmt.Println("\n Задача 2")
    nums := []int{2, 4, 6, 8, 10}

    firstWay(nums)  // Первый способ: использование небуферизованного канала и WaitGroup
    secondWay(nums) // Второй способ: использование буферизированного канала
    thirdWay(&nums, 2) // Третий способ: разделение массива на сегменты и вычисления in-place.
}

func firstWay(nums []int) {
    fmt.Println("   Первый способ: использование канала и WaitGroup.")
    squareChan := make(chan int)

    var wg sync.WaitGroup

    for _, num := range nums {
        wg.Add(1)
        go func(x int) {
            defer wg.Done()
            square := x * x
            fmt.Println("Calculated: ", square)
            squareChan <- square
        }(num)
    }

    // Закрываем канал после завершения всех горутин
    go func() {
        wg.Wait()
        close(squareChan) 
    }()

    for square := range squareChan {
        fmt.Println("Read: ", square)
    }
}

func secondWay(nums []int) {
    fmt.Println("   Второй способ: использование буферизированного канала")
    squareChan := make(chan int, len(nums))

    for _, num := range nums {
        go func(x int) {
            square := x * x
            fmt.Println("Calculated: ", square)
            squareChan <- square // Операция будет блокирована, до тех пор пока в канале есть другое значение 
        }(num)
    }

    for i := 0; i < len(nums); i++ {
        square := <-squareChan // Операция будет ожидать появления значений в канале.
        // Поскольку канал небуферизованный главная горутина будет блокироваться до тех пор,
        // пока одна из горутин не отправит в канал значение
        fmt.Println("Read: ", square)
    }
}

func thirdWay(nums *[]int, segments int) {
    fmt.Println("   Третий способ: сегментация массива и синхронизация через WaitGroup")
    n := len(*nums)
    seg := 0
    // Вычисляем размер сегмента 
    if segments >= n {
        // Если количество сегментов больше или равно количеству элементов, каждый сегмент будет содержать по одному элементу.
        seg = 1
    } else if segments < 2 { // Если указано количество сегментов меньше 2, то будет только один сегмент для обработки.
        seg = n
    } else {
        // Иначе равномерно распределяем элементы
        seg = n / segments
        // Если заданного количества сегментов не хватит, например как в данном примере использования:
        // seg = 5 / 2 = 2, то программа выделит еще одну горутину для обработки остатка массива.
    }

    var wg sync.WaitGroup
    // Гонки данных не возникнет - каждая горутина имеет отдельный сегмент
    for i := 0; i < n; i += seg {
        wg.Add(1)
        go func(a, b int) { 
            defer wg.Done()
            // Вычисляем квадраты для элементов в интервале [a, b)
            for j := a; j < b && j < n; j++ {
                x := (*nums)[j]
                sq := x * x
                fmt.Println("Calculated: ", sq)
                (*nums)[j] = sq 
                fmt.Println("Inserted: ", sq)
            }
        }(i, i + seg)
    }

    wg.Wait()

    fmt.Println("Убедимся, что порядок сохранился: ", (*nums))
}

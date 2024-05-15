package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
    nums := []int{2, 4, 6, 8, 10}
    fmt.Println("\n Задача 3")

    firstWay(nums) // Первый способ: использование буферизованного канала
    secondWay(nums, 2) // Второй способ: сегментация массива и использование общей переменной с sync.Mutex
    thirdWay(nums, 2) // Третий способ: сегментация массива и использование общей переменной с атомарными операциями.
}

func firstWay(nums []int) {
    fmt.Println("   Первый способ: использование буферизованного канала.")
    ans := 0
    ansChan := make(chan int, len(nums)) // Создаем буферизованный канал с длинной len(nums)

    for _, num := range nums {
        go func(x int) {
            square := x * x
            fmt.Println("Calculated: ", square)
            ansChan <- square
        }(num)
    }

    // Считываем все квадраты и добавляем их к ответу
    for i := 0; i < len(nums); i++ { 
        square := <- ansChan
        ans += square
        fmt.Println("Added : ", square)
    }
    fmt.Println("Sum: ", ans)
}

func secondWay(nums []int, segments int) {
    fmt.Println("   Второй способ: сегментация массива и использование общей переменной с Mutex.")
    n := len(nums)
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

    ans := 0 // В этом коде имеется доступ к общей переменной из нескольких горутин,
    var mx sync.Mutex // поэтому используется мутекс, чтобы избежать гонки данных.
    var wg sync.WaitGroup
    for i := 0; i < n; i += seg {
        wg.Add(1)
        go func(a, b int) { 
            defer wg.Done()
            // Вычисляем квадраты для элементов в интервале [a, b)
            for j := a; j < b && j < n; j++ {
                x := nums[j]
                sq := x * x
                fmt.Println("Calculated: ", sq)
                mx.Lock()
                ans += sq
                mx.Unlock()
                fmt.Println("Added: ", sq)
            }
        }(i, i + seg)
    }

    wg.Wait()
    fmt.Println("Sum: ", ans)
}

func thirdWay(nums []int, segments int) {
    fmt.Println("   Третий способ: сегментация массива и использование общей переменной с атомарными операциями.")
    n := len(nums)
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

    var ans int32 = 0 // В этом коде имеется доступ к общей переменной из нескольких горутин.
    var wg sync.WaitGroup
    for i := 0; i < n; i += seg {
        wg.Add(1)
        go func(a, b int) { 
            defer wg.Done()
            // Вычисляем квадраты для элементов в интервале [a, b)
            for j := a; j < b && j < n; j++ {
                x := nums[j]
                sq := x * x
                fmt.Println("Calculated: ", sq)
                atomic.AddInt32(&ans, int32(sq)) // Использование атомарной операции для предотвращения гонки данных.
                // Данные операция обеспечивает что,
                // планировщик GO не будет выполнять никакую другую горутину во время прибавления.
                fmt.Println("Added: ", sq)
            }
        }(i, i + seg)
    }

    wg.Wait()
    fmt.Println("Sum: ", ans)
}

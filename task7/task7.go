package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
    fmt.Println("\n Задача 7")
    firstWay()
    time.Sleep(2 * time.Second)
    secondWay()
}

// С помощью RWMutex
var mx sync.RWMutex

func firstWay() {
    fmt.Println("   Первый способ: использование RWMutex")
    countMap := make(map[int]int)

    var wg sync.WaitGroup
    for i := 0; i < 3; i++ {
        wg.Add(1)
        go mapWriter(&countMap, &wg, i+1)
    }

    wg.Wait()
    for k, v := range countMap {
		fmt.Printf("Key: %v, Value: %v\n", k, v)
    }
}

// Каждый из записывателей, будет пытаться инкрементировать значения в [0, 9] по порядку.
func mapWriter(countMap *map[int]int, wg *sync.WaitGroup, writerID int) {
    defer wg.Done()
    for i := 0; i < 10; i++ {
        fmt.Printf("Записыватель %d пытается инкрементировать по ключу %d\n", writerID, i)
        // Когда вызывается mx.Lock(), гарантируется, что только одна горутина получит доступ к мапе
        // остальные будут ждать вызова mx.Unlock(), следовательно гонки данных не возникнет.
        mx.Lock()
        (*countMap)[i]++
        fmt.Printf("Записыватель %d инкрементирует по ключу %d\n", writerID, i)
        time.Sleep(250 * time.Millisecond)
        mx.Unlock()
    }
}

func secondWay() {
    fmt.Println("   Второй способ: использование sync.Map")
    var countMap sync.Map
    var wg sync.WaitGroup
    for i := 0; i < 3; i++ {
        wg.Add(1)
        go syncMapWriter(&countMap, &wg, i+1)
    }
    wg.Wait()
    countMap.Range(func(key, value interface{}) bool {
		fmt.Printf("Key: %v, Value: %v\n", key, value)
		return true
	})
}

// Каждый из записывателей, будет пытаться записать значение из longCalculations() по ключам (0..=9).
func syncMapWriter(countMap *sync.Map, wg *sync.WaitGroup, writerID int) {
    defer wg.Done()
    for i := 0; i < 10; i++ {
        fmt.Printf("Записыватель %d пытается записать по ключу %d\n", writerID, i)
        // sync.Map гарантирует безопасную конкурентую запись, но не рекомендуется к использованию
        // потому что может привести к потере типовой безопасности т.к. использует interface{},
        // сложнее в использовании, имеет ограниченный набор методов
        // и в большинстве случаев менее производительна чем обычная map + sync.RWMutex.
        countMap.Store(i, longCalculations())
        fmt.Printf("Записыватель %d записал по ключу %d\n", writerID, i)
    }
}

// Имитируем длительные вычисления
func longCalculations() int {
    time.Sleep(250 * time.Millisecond)
    return 69
}

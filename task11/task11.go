package main

import "fmt"


// В качестве неупорядоченного множества будем использовать map[T]bool
// т.к. map хранит значения без определенного порядка.
//
// T comparable - это все типы которые поддерживают операции == и != 
// map - требует чтобы тип ключа реализовывал comparable.
type Set[T comparable] map[T]bool

func main() {
    fmt.Println("\n Задача 11")
    var set1 Set[int] = make(Set[int])
    var set2 Set[int] = make(Set[int])
    for i := 0; i < 5; i++ {
        set1[i] = true
    }
    for i := 3; i < 10; i++ {
        set2[i] = true
    }
    set3 := intersection(set1, set2)
    fmt.Printf("Множество 1: %v\nМножество 2: %v\nПересечение: %v\n", set1, set2, set3)
}

func intersection[T comparable](set1 Set[T], set2 Set[T]) Set[T] {
    newSet := make(Set[T])
    for k, v := range set1 {
        // Ищем только значения содержащиеся в обоих множествах
        if v && set2[k] {
            newSet[k] = true
        }
    }
    return newSet
}

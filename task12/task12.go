package main

import "fmt"


type Set[T comparable] map[T]bool
// Само множество и пустое множество называют несобственными подмножествами,
// остальные подмножества называют собственными.
func main() {
    fmt.Println("\n Задача 12")
    someStrings := []string{"cat", "cat", "dog", "cat", "tree"}
    set := make(Set[string])
    for _, s := range someStrings {
        set[s] = true
    }
    properSets := allProperSets(set)
    fmt.Printf("Исходное множество: %v\n", set)
    fmt.Println("Собственные подмножества:")
    for _, s := range properSets {
        fmt.Printf("%v\n", s)
    }
}

func allProperSets[T comparable](set Set[T]) []Set[T] {
    var properSets []Set[T]
    keys := make([]T, 0, len(set))
    for k, v := range set {
        if v {
            keys = append(keys, k)
        }
    }
    total := 1 << len(keys) // всего подмножеств 2^n, собственных выйдет 2^n - 2 (минус пустое и исходное)
    for i := 1; i < total; i++ { 
        subset := make(Set[T])
        for j, k := range keys {
            if i&(1<<j) != 0 { // проверяем, установлен ли j-й бит в числе i, если установлен значит
                // включаем элемент в подмножество, таким образом мы составим все собственные подмножества,
                // поскольку переберем все возможные комбинации среди 2^n (кроме пустого и исходного множеств).
                // Пример:
                // i = 0b000101 => мы должны включить первый и третий ключи (0 и 2 индексы)
                subset[k] = true
            }
        }
        if len(subset) != len(set) { // собственное подмножество не может быть изначальным множеством
            properSets = append(properSets, subset)
        }
    }
    return properSets
}

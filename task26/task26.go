package main

import (
	"fmt"
	"strings"
)


type Set[T comparable] map[T]bool

func main() {
    fmt.Println("\n Задача 26")
    fmt.Println("abcd: ", isAllCharsUnique("abcd"))
    fmt.Println("abCdefAaf: ", isAllCharsUnique("abCdefAaf"))
    fmt.Println("aabcd: ", isAllCharsUnique("aabcd"))
}

// Определяет содержит ли строка только уникальные символы, с помощью неупорядоченного множества
func isAllCharsUnique(s string) bool {
    lowerS := strings.ToLower(s)
    set := make(Set[byte])
    
    for i := 0; i < len(s); i++ {
        // Если такой элемент уже есть значит он не уникален
        if set[lowerS[i]] {
            return false
        } else {
        // Отмечаем новый элемент
            set[lowerS[i]] = true
        }
    }
    return true
}

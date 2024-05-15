package main

import "fmt"

// Подразумевая, что мы создаем огромную строку 'v' длиной 2^11,
// проблема заключается в том что переменная justString, не создает новую строку длиной 100 символов,
// а ссылается на огромную строку 'v', этот код может привести к утечке памяти (удержанию слишком большого объема).
//
// var justString string
// func someFunc() {
//   v := createHugeString(1 << 10)
//   justString = v[:100]
// }
//
// func main() {
//   someFunc()
// }

func main() {
    fmt.Println("\n Задача 15")
    fmt.Println("Посмотри на код!")
}

var justString string

func someFunc() {
    v := createHugeString(1 << 10)
    res := []byte{}
    // Вместо того чтобы ссылаться на 'v' создадим новую строку со 100 символами из 'v'.
    for i := 0; i < 100; i++ {
        res = append(res, v[i])
    }
    justString = string(res)

    // Либо можно воспользоваться copy()
    copiedString := make([]byte, 100)
    copy(copiedString, v[:100])
    justString = string(copiedString)
}

func createHugeString(length int) string {
    res := []byte{}
    for i := 0; i < length; i++ {
        res = append(res, byte('x'))
    }
    return string(res)
}

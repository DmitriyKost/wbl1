package main

import "fmt"


func main() {
    fmt.Println("\n Задача 14")
    var (
        num int = 42
        str string = "Hello"
        boolean bool = true
        ch chan interface{} = make(chan interface{})
    )

    var iface interface{}
    printType(iface)

    iface = num
    printType(iface)

    iface = str
    printType(iface)

    iface = boolean
    printType(iface)

    iface = ch
    printType(iface)
}

// Проверяем и выводим тип переменной содержащейся в iface.
func printType(iface interface{}) {
    switch iface.(type) {
    case int:
        fmt.Println("Переменная является int.")
    case string:
        fmt.Println("Переменная является string.")
    case bool:
        fmt.Println("Переменная является bool.")
    case chan interface{}:
        fmt.Println("Переменная является channel.")
    default:
        fmt.Println("Неизвестный тип переменной.")
    }
}

package main

import "fmt"


func main() {
    fmt.Println("\n Задача 13")
    var a, b int = 69, 420
    // 'a, b = b, a' не создает временную переменную явно. 
    // Он создает временный кортеж, содержащий значения 'a' и 'b'. 
    // Когда 'b' присваивается 'a', а 'a' присваивается 'b', они напрямую обмениваются значениями 
    // без необходимости использования отдельной временной переменной. 
    //
    // Однако я не уверен подходит ли это под условие задачи...
    fmt.Println("Меняем местами с помощью: a, b = b, a")
    fmt.Printf("a = %d, b = %d\n", a, b)
    a, b = b, a
    fmt.Printf("a = %d, b = %d\n", a, b)

    fmt.Println("Меняем местами с помощью: XOR")
    fmt.Printf("a = %d, b = %d\n", a, b)
    a ^= b
    b ^= a
    a ^= b   
    fmt.Printf("a = %d, b = %d\n", a, b)

    fmt.Println("Меняем местами с помощью: арифметических операций")
    fmt.Printf("a = %d, b = %d\n", a, b)
    a += b
    b = a - b
    a = a - b
    fmt.Printf("a = %d, b = %d\n", a, b)
}
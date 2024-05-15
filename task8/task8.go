package main

import "fmt"


func main() {
    fmt.Println("\n Задача 8")
    var variable int64 = 12
    // Устанавливаем третий бит в ноль (программа должна вывести 4)
    firstWay(&variable, 3, true)
    // Устанавливаем четвертый бит в единицу (программа должна вывести 20)
    firstWay(&variable, 4, false)
    variable = 12
    // Устанавливаем третий бит в ноль (программа должна вывести 4)
    secondWay(&variable, 3, true)
    // Устанавливаем четвертый бит в единицу (программа должна вывести 20
    secondWay(&variable, 4, false)
}

func firstWay(variable *int64, i int, toZero bool) {
    fmt.Println("   Первый способ")
    fmt.Println("До: ", int(*variable))
    if toZero {
        *variable &^= 1 << i // Сдвигаем 1 на i битов, и выполняем операцию AND NOT между variable и битом на i-й позиции
        // Пример:
        // variable: 0b0110
        //  3-й бит: 0b0100 
        // 0b0110 & NOT(0b0100) = 0b0110 & 1b1011 = 0b0010
        // т.е. мы установим в 0 только i-й бит, независимо от значения этого бита в variable
    } else {
        *variable |= 1 << i // Сдвигаем 1 на i битов, и выполняем операцию OR между variable и битом на i-й позиции
    }
    fmt.Println("После: ", int(*variable))
}

func secondWay(variable *int64, i int, toZero bool) {
    fmt.Println("   Второй способ")
    fmt.Println("До: ", int(*variable))
    if toZero {
        *variable -= *variable & (1 << i)
        // *variable & (1 << i) - проверяем если на i-й позиции уже 0 
        // в таком случае *variable & (1 << i) = 0 
        // иначе = 1 
        // таким образом мы установим i-й бит в ноль.
    } else {
        *variable += (1 << i) - (*variable & (1 << i))
        // (1 << i) - сдвигаем бит на i-ю позицию
        // (*variable & (1 << i)) - смотрим, установлен ли бит в 1, если да то
        // (1 << i) - (*variable & (1 << i)) = 0 иначе = 1
        // таким образом мы установим i-й бит в единицу.
    }
    fmt.Println("После: ", int(*variable))
}

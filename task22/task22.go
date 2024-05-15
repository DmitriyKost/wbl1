package main

import (
	"fmt"
	"math"
)


func main() {
    fmt.Println("\n Задача 22")
    var a, b float64 = 2 << 23, 2 << 24
    // В качестве типа данных для подобной задачи можно также использовать int64,
    // Однако float64 гораздо больше, и мы с меньшей вероятностью выйдем за допустимые пределы.
    performActions(a, b)
}

func performActions(a, b float64) {
    // Перемножение
	product := a * b
	fmt.Printf("Произведение %f и %f: %f\n", a, b, product)

	// Деление
	division := a / b
	fmt.Printf("Деление %f на %f: %f\n", a, b, division)

	// Сложение
	sum := a + b
	fmt.Printf("Сумма %f и %f: %f\n", a, b, sum)

	// Вычитание
	subtraction := a - b
	fmt.Printf("Разность %f и %f: %f\n", a, b, subtraction)

	// Проверка на переполнение типа float64, (крайне маловероятно, максимальное значение примерно 1.7 * 10^308)
	if math.IsInf(product, 0) || math.IsInf(division, 0) || math.IsInf(sum, 0) || math.IsInf(subtraction, 0) {
		fmt.Println("Результаты вышли за пределы допустимых значений.")
	} else {
		fmt.Println("Результаты в допустимых пределах.")
	}
}

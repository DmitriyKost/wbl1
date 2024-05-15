package main

import "fmt"

func main() {
    fmt.Println("\n Задача 17")
    nums := []int{1, 2, 3, 37, 69, 420}
    fmt.Println("Наш массив: ", nums)
    fmt.Printf("Ищем 69... Нашли idx = %d!\n", BinarySearch(&nums, 69))
    fmt.Printf("Ищем 96... Не нашли idx = %d :(\n", BinarySearch(&nums, 96))
}

type Number interface {
    int | int8 | int16 | int32 | int64 |
    float32 | float64
}

// Функция ищет target в указанном массиве (подразумевая, что он отсортирован в неубывающем порядке),
// Возвращает индекс найденного значения или -1 если не обнаружит.
func BinarySearch[T Number](nums *[]T, target T) int {
    left, right := 0, len(*nums) - 1
    for left <= right {
        mid := (left + right) / 2
        // Возвращаем индекс если обнаружили target
        if (*nums)[mid] == target {
            return mid
        // Если nums[mid] меньше target значит nums[:left+1] нас не интересуют
        } else if (*nums)[mid] < target {
            left = mid + 1
        // Если nums[mid] больше target значит nums[right:] нас не интересуют
        } else {
            right = mid - 1
        }
    }
    return -1 
}

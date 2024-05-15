package main

import "fmt"

type Number interface {
    int | int8 | int16 | int32 | int64 |
    float32 | float64
}


func main() {
    fmt.Println("\n Задача 23")
    nums := []int{1, 2, 37, 69, 420}
    fmt.Println("Слайс: ", nums)
    if err, ans := deleteAt(&nums, 3); err == nil {
        fmt.Println("Удаленный элемент: ", ans)
    } else {
        fmt.Println("Ошибка: ", err)
        return
    }
    fmt.Println("Слайс после удаления: ", nums)
}

func deleteAt[T Number](nums *[]T, i int) (error, T) {
    if len(*nums) - 1 < i {
        return fmt.Errorf("Index is %d while len is %d!", i, len(*nums)), -1
    } else if i < 0 {
        return fmt.Errorf("Index is %d < 0!", i), -1
    }
    ans := (*nums)[i]
    *nums = append((*nums)[:i], (*nums)[i+1:]...)
    return nil, ans
}

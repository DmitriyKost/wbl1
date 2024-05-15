package main

import (
	"fmt"
	"sync"
)

type Number interface {
    int | int8 | int16 | int32 | int64 |
    float32 | float64
}

func main() {
    fmt.Println("\n Задача 16")
    nums := []int{37, 69, 420, 2004, 4, 21}
    fmt.Println("До сортировки: ", nums)
    QuickSort(&nums)
    fmt.Println("После :", nums)
}

// Функция сортирует массив чисел in-place, используя параллельные вычисления
func QuickSort[T Number](arr *[]T) {
    var wg sync.WaitGroup
    wg.Add(1)
    quickSortPrivate(arr, 0, len(*arr)-1, &wg)
    wg.Wait()
}

// Возвращает индекс медианы трех значений arr[low], arr[mid], arr[high].
// Это используется для выбора опорного элемента при разделении массива на две части в функции partition().
//
// В среднем такой выбор опорного элемента приводит к лучшей ассимптотике.
func medianOfThree[T Number](arr []T, low, high int) int {
	mid := low + (high-low)/2

	if arr[low] < arr[mid] {
		if arr[mid] < arr[high] {
			return mid
		} else if arr[low] < arr[high] {
			return high
		} else {
			return low
		}
	} else {
		if arr[low] < arr[high] {
			return low
		} else if arr[mid] < arr[high] {
			return high
		} else {
			return mid
		}
	}
}

// Разделяет массив на две части относительно опорного элемента.
// Опорный элемент выбирается с помощью функции medianOfThree, затем массив перестраивается таким образом,
// что все элементы, меньшие опорного, находятся слева от него, а все элементы, большие опорного, справа от него.
func partition[T Number](arr *[]T, low, high int) int {
	pivotIndex := medianOfThree((*arr), low, high)
	(*arr)[pivotIndex], (*arr)[low] = (*arr)[low], (*arr)[pivotIndex]

	pivot := (*arr)[low]
	i, j := low-1, high+1

	for {
		for {
			i++
			if (*arr)[i] >= pivot {
				break
			}
		}
		for {
			j--
			if (*arr)[j] <= pivot {
				break
			}
		}
		if i >= j {
			return j
		}
		(*arr)[i], (*arr)[j] = (*arr)[j], (*arr)[i]
	}
}

// Выполняет рекурсивный процесс сортировки. 
// Вызывает partition для разделения массива на две части, а затем рекурсивно сортирует каждую из них.
func quickSortPrivate[T Number](arr *[]T, low, high int, wg *sync.WaitGroup) {
    defer wg.Done()
	if low < high {
		pi := partition(arr, low, high)
		var newWg sync.WaitGroup
		newWg.Add(2)
		go quickSortPrivate(arr, low, pi, &newWg)
		go quickSortPrivate(arr, pi+1, high, &newWg)
		newWg.Wait()
	}
}

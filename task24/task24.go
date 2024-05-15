package main

import (
	"fmt"
	"math"
)


func main() {
    fmt.Println("\n Задача 24")
    p1 := NewPoint(1, 1)
    p2 := NewPoint(37, 69)
    fmt.Println("Точка 1:")
    fmt.Println(p1.GetCoordinates())
    fmt.Println("Точка 2:")
    fmt.Println(p2.GetCoordinates())
    fmt.Println("Расстояние между точками ", p1.Distance(p2))
}

type Number interface {
    int | int8 | int16 | int32 | int64 |
    float32 | float64
}

type Point[T Number] struct {
    x, y T
}

func NewPoint[T Number](x, y T) Point[T] {
    return Point[T]{x, y}
}

func (p *Point[T]) Distance(p2 Point[T]) float64 {
    x2, y2 := p2.GetCoordinates()
    x := abs(p.x - x2)
    y := abs(p.y - y2)
    return math.Sqrt(float64(
        x*x + y*y,
    ))
}

func (p *Point[T]) GetCoordinates() (T, T) {
    return p.x, p.y
}

func abs[T Number](a T) T {
    if a < 0 {
        return -1 * a
    }
    return a
}

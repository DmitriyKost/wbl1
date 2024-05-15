package main

import "fmt"

// Структуры с публичными полями
type Human struct {
    Name string
    Age uint
}

func (h *Human) Introduce() {
    fmt.Printf("Hi, my name is %s, I'm %d years old!\n", h.Name, h.Age)
}

func (h *Human) SayBye() {
    fmt.Println("Bye-bye!")
}

func NewHuman(name string, age uint) Human {
    return Human{Name: name, Age: age}
}

// Структура Action 'наследует' поля и методы 'родительской' структуры Human, с помощью 'Анонимного поля',
// Анонимность поля заключается в том, что при вызове методов, поле не нужно указывать:
//
//  action := NewAction("Ivan", 18)
//  action.Introduce()
//  action.SayBye()
// 
type Action struct {
    Human
}

func NewAction(name string, age uint) Action {
    return Action{Human: NewHuman(name, age)}
}

// Структуры с приватными полями
type HumanPrivate struct {
    name string
    age uint
}

func (h *HumanPrivate) Introduce() {
    fmt.Printf("Hi, my name is %s, I'm %d years old, but it's a secret!\n", h.name, h.age)
}

func (h *HumanPrivate) SayBye() {
    fmt.Println("Bye-bye!")
}

func NewHumanPrivate(name string, age uint) HumanPrivate {
    return HumanPrivate{name: name, age: age}
}

// Структура ActionPrivate 'наследует' поля и методы 'родительской' структуры HumanPrivate, с помощью 'Анонимного поля',
// Анонимность поля заключается в том, что при вызове методов, поле не нужно указывать:
//
//  actionPrivate := NewActionPrivate("Ivan", 18)
//  actionPrivate.Introduce()
//  actionPrivate.SayBye()
// 
type ActionPrivate struct {
    HumanPrivate
}

func NewActionPrivate(name string, age uint) ActionPrivate {
    return ActionPrivate{HumanPrivate: NewHumanPrivate(name, age)}
}

func main() {
    fmt.Println("\n Задача 1")
    // Пример с публичными полями
    fmt.Println("Структура Human с публичными полями, переданными в структуру.")
    youngLady := Human{Name: "Julia", Age: 25}
    youngLady.Introduce()
    youngLady.SayBye()
    fmt.Println("Структура Human с публичными полями, переданными в функцию 'конструктор'.")
    youngLady = NewHuman("Malena", 19)
    youngLady.Introduce()
    youngLady.SayBye()
    // Можно создать структуру Action передав уже существующую структуру Human, чтобы скопировать ее поля и изменить их.
    fmt.Println("Структура Action с анонимным полем Human - существующей переменной, 'унаследованные' публичные поля были изменены.")
    action := Action{Human: youngLady}
    action.Name = "John"
    action.Age = 28
    action.Introduce()
    action.SayBye()
    fmt.Println("Структура Action с полями, переданными в функцию 'конструктор'.")
    // Также можно создать структуру с помощью функции NewAction.
    action = NewAction("Bob", 18) 
    action.Introduce()
    action.SayBye()


    // Пример с приватными полями
    // Аналогично примеру выше, уже существующую структуру HumanPrivate можно передать в структуру ActionPrivate,
    // но поля изменить уже нельзя - они приватные.
    fmt.Println("Структура HumanPrivate с приватными полями, переданными в структуру.")
    mysteriousLady := HumanPrivate{name: "Karina", age: 21}
    mysteriousLady.Introduce()
    mysteriousLady.SayBye()
    fmt.Println("Структура ActionPrivate с анонимным полем HumanPrivate - существующей переменной, 'унаследованные' приватные поля нельзя изменить.")
    mysteriousLadyAction := ActionPrivate{HumanPrivate: mysteriousLady}
    mysteriousLadyAction.Introduce()
    mysteriousLadyAction.SayBye()
    // Также можно использовать функции NewActionPrivate и NewHumanPrivate.
    fmt.Println("Структура HumanPrivate с полями, переданными в функцию 'конструктор'.")
    mysteriousMan := NewHumanPrivate("Dmitriy", 20)
    mysteriousMan.Introduce()
    mysteriousMan.SayBye()
    fmt.Println("Структура ActionPrivate с полями, переданными в функцию 'конструктор'.")
    mysteriousManAction := NewActionPrivate("Val", 20)
    mysteriousManAction.Introduce()
    mysteriousManAction.SayBye()
}

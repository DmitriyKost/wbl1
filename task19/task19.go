package main

import "fmt"



func main() {
    fmt.Println("\n Задача 19")
    s := "Hello, World!"
    fmt.Println("Исходная строка: ", s)
    fmt.Println("   Первый способ: новая строка")
    fmt.Println(reversedString(s))

    fmt.Println("   Второй способ: in-place")
    reverseString(&s)
    fmt.Println(s)

    fmt.Println("   Третий способ: с Кириллицей")
    fmt.Println(reverseCyrillic("Привет, Мир!"))
}

// Возвращает новую перевернутую строку
func reversedString(s string) string {
    ans := []byte{}
    for i := 0; i < len(s); i++ {
        ans = append(ans, s[len(s)-i-1])
    }
    return string(ans)
}

// Переворачивает строку 'in-place' (в Go строки иммутабельны, поэтому решение выглядит так).
func reverseString(s *string) {
    ans := []byte{}
    for i := 0; i < len(*s); i++ {
        ans = append(ans, (*s)[len(*s)-i-1])
    }
    *s = string(ans)
}

// Поскольку функции выше работают с byte, перевернутые строки на русском отображаются некорректно.
//
// Эта функция использует rune.
func reverseCyrillic(s string) string {
    ans := make([]rune, len(s))
    for i, r := range s {
        ans[len(s)-i-1] = r
    }
    return string(ans)
}

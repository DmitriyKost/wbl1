package main

import "fmt"


func main() {
    fmt.Println("\n Задача 20")
    fmt.Println(reverseWords("snow dog sun"))
}

// s может содержать начальные или конечные пробелы или несколько пробелов между словами.
// Возвращаемая строка будет содержать только по одному пробелу между словами.
func reverseWords(s string) string {
    if len(s) == 1 && s[0] != ' ' { return s }
    ans := ""
    j, k := len(s) - 1, len(s) - 1
    for j > 0 {
        // Пропускаем пробелы
        for j > 0 && s[j] == ' ' {
            j--
        }
        // Если дошли до конца выходим из цикла
        if j == 0 && s[j] == ' ' { break }
        k = j
        // Ищем конец слова
        for j > 0 && s[j-1] != ' ' {
            j--
        }
        // Добавляем новое слово
        ans += s[j:k+1] + " "
        j--
    }
    // Возвращаем перевернутую строку, без последнего пробела
    return ans[0:len(ans)-1]
}

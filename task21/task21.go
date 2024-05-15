package main

import (
	"fmt"
)

type Reader interface {
	Read() string
}

// Структура FileReader, реализующая интерфейс Reader
type FileReader struct{}

func (fr *FileReader) Read() string {
	return "Читаем из файла"
}

type Writer interface {
	Write(string)
}

// Адаптер, преобразующий Reader в Writer
type ReaderToWriterAdapter struct {
	reader Reader
}

func (rw *ReaderToWriterAdapter) Write(data string) {
	fmt.Println("Адаптируем")
    fmt.Println("Адаптировано: ", rw.reader.Read())
}

func main() {
    fmt.Println("\n Задача 21")
	// Создаем экземпляр FileReader
	fileReader := &FileReader{}

	// Создаем адаптер, передавая FileReader в качестве Reader
	adapter := &ReaderToWriterAdapter{reader: fileReader}

	// Вызываем метод Write у адаптера, который адаптирует чтение к записи
	adapter.Write("Data to write")
}

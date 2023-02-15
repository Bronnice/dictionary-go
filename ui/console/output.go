package console

import (
	"dictionary-go/dictionary"
	"fmt"
)

// Вывод строки в консоль
func Println(a ...any) {
	fmt.Println(a...)
}

// Форматированный вывод словаря в консоль
func PrintDictionary(dictionary *dictionary.Dictionary) {
	Println("Список слов:")
	for word, translate := range dictionary.WordMap() {
		Printf("%s - %s", word, translate)
		Println()
	}
}

// Форматированный вывод строки в консоль
func Printf(str string, a ...any) {
	fmt.Printf(str, a...)
	Println()
}

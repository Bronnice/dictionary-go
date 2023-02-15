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
		Println(Sprintf("%s - %s", word, translate))
	}
}

// Форматированный вывод строки в консоль
func Printf(str string, a ...any) {
	fmt.Printf(str, a...)
}

func Sprintf(format string, a ...any) string {
	return fmt.Sprintf(format, a...)
}

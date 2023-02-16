package console

import (
	"dictionary-go/dictionary"
	"fmt"
)

// Вывод строки в консоль
func Println(str ...any) {
	fmt.Println(str...)
}

// Форматированный вывод словаря в консоль
func PrintDictionary(dictionary *dictionary.Dictionary) {
	Println("Список слов:")
	for word, translate := range dictionary.WordMap() {
		PrintlnFormatted("%s - %s", word, translate)
	}
}

// Форматирование строки
func PrintlnFormatted(format string, a ...any) {
	Println(fmt.Sprintf(format, a...))
}

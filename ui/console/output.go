package console

import (
	"dictionary-go/dictionary"
	"fmt"
)

// Вывод строки в консоль
func Println(str ...string) {
	fmt.Println(str)
}

// Форматированный вывод словаря в консоль
func PrintDictionary(dictionary *dictionary.Dictionary) {
	Println("Список слов:")
	for word, translate := range dictionary.WordMap() {
		Println(fmt.Sprintf("%s - %s", word, translate))
	}
}

// Форматирование строки
func PrintlnFormatted(format string, str ...string) {
	Println(fmt.Sprintf(format, str))
}

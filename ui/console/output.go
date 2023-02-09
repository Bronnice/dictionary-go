package console

import (
	"dictionary-go/dictionary"
	"fmt"
)

//Вывод строки в консоль
func Println(str string) {
	fmt.Println(str)
}

//Форматированный вывод словаря в консоль
func PrintDictionary(dictionary *dictionary.Dictionary) {
	Println("Список слов:")
	for word, translate := range dictionary.WordMap() {
		fmt.Printf("%s - %s\n", word, translate)
	}
}

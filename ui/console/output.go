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
	fmt.Println("Номер словаря: ", dictionary.Number)
	Println("Список слов:")
	for word, translate := range dictionary.WordMap() {
		Println(word + " - " + translate)
	}
}

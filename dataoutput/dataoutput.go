package dataoutput

import (
	"dictionary-go/dictionary"
	"fmt"
)

//Форматированный вывод словаря в консоль
func DictionaryOutput(dictionary dictionary.Dictionary) {
	fmt.Println("Имя словаря: " + dictionary.Name)
	fmt.Println("Список слов:")
	for word, translate := range dictionary.WordMap() {
		fmt.Println(word + " - " + translate)
	}
}

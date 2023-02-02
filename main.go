package main

import (
	"dictionary-go/dataoutput"
	"dictionary-go/dictionary"
)

func main() {
	//Пример просмотра словаря
	exampleDictionary := dictionary.NewDictionary("exampleName")

	dataoutput.DictionaryOutput(*exampleDictionary)
}

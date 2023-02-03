package main

import (
	"dictionary-go/dataoutput"
	"dictionary-go/datasource"
	"dictionary-go/dictionary"
)

func main() {
	//Пример инициалиации словаря
	exampleDictionary := dictionary.NewDictionary("exampleName")

	//Пример добавление слова в словарь
	datasource.AddWord(exampleDictionary, "name", "1234")
	//Пример просмотора словаря
	dataoutput.DictionaryOutput(*exampleDictionary)
}

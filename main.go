package main

import (
	"dictionary-go/dataoutput"
	"dictionary-go/datasource"
	"dictionary-go/dictionary"
)

func main() {
	//Пример просмотра словаря
	exampleDictionary := dictionary.NewDictionary("exampleName")

	datasource.AddNewIssue(*exampleDictionary, "name", "1234")
	dataoutput.DictionaryOutput(*exampleDictionary)
}

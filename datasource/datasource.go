package datasource

import (
	"dictionary-go/dictionary"
)

//Создает новую запись в словарь
func AddNewIssue(dictionary *dictionary.Dictionary, word, translate string) {
	dictionary.WordMap()[word] = translate
}

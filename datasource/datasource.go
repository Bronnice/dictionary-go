package datasource

import (
	"dictionary-go/dictionary"
)

//Создает новую запись в словарь
func AddWord(dictionary *dictionary.Dictionary, word, translate string) {
	dictionary.WordMap()[word] = translate
}

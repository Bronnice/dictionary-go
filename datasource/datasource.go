package datasource

import (
	"dictionary-go/dictionary"
)

//Добавляет новое слово и его значение  в словарь
func AddWord(dictionary *dictionary.Dictionary, word, translate string) {
	dictionary.WordMap()[word] = translate
}

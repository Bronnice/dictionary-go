package datasource

import (
	"dictionary-go/dictionary"
	"dictionary-go/validator"
	"strings"
)

var dictionaries []dictionary.Dictionary = []dictionary.Dictionary{
	*dictionary.NewDictionary("English", validator.NewEnglishWordValidator()),
	*dictionary.NewDictionary("Numbers", validator.NewNumberWordValidator()),
}

// Находит словарь по имени и возвращает на него ссылку
func GetDictionaryByName(name string) *dictionary.Dictionary {
	for _, dictionary := range dictionaries {
		if strings.EqualFold(name, dictionary.Name()) {
			return &dictionary
		}
	}
	return nil
}

// Возвращает все имена доступных словарей
func GetDictionariesNames() [2]string {
	var str [2]string
	for i, dictionary := range dictionaries {
		str[i] = dictionary.Name()
	}
	return str
}

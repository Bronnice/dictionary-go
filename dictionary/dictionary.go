package dictionary

import (
	"dictionary-go/validation"
)

//Структура для словаря
type Dictionary struct {
	Number     int
	wordMap    map[string]string
	validation validation.Validator
}

//Конструктор для Dictionary
func NewDictionary(validation validation.Validator) *Dictionary {
	return &Dictionary{
		Number:     -1,
		wordMap:    make(map[string]string),
		validation: validation,
	}
}

//Получить список слов
func (dictionary *Dictionary) WordMap() *map[string]string {
	return &dictionary.wordMap
}

func (dictionary *Dictionary) Validation() *validation.Validator {
	return &dictionary.validation
}

//Добавляет новое слово и его значение  в словарь
func (dictionary *Dictionary) AddWord(word, translate *string) {
	dictionary.wordMap[*word] = *translate
}

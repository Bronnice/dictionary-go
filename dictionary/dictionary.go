package dictionary

import (
	"dictionary-go/validator"
)

//Структура для словаря
type Dictionary struct {
	Number    int
	wordMap   map[string]string
	Validator validator.Validator
}

//Конструктор для Dictionary
func NewDictionary(validator validator.Validator) *Dictionary {
	return &Dictionary{
		Number:    -1,
		wordMap:   make(map[string]string),
		Validator: validator,
	}
}

//Получить список слов
func (dictionary *Dictionary) WordMap() map[string]string {
	return dictionary.wordMap
}

//Добавляет новое слово и его значение  в словарь
func (dictionary *Dictionary) AddWord(word, translate *string) {
	dictionary.wordMap[*word] = *translate
}

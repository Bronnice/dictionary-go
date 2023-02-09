package dictionary

import (
	"dictionary-go/validator"
	"errors"
)

//Структура для словаря
type Dictionary struct {
	wordMap   map[string]string
	Validator validator.Validator
}

//Конструктор для Dictionary
func NewDictionary(validator validator.Validator) *Dictionary {
	return &Dictionary{
		wordMap:   make(map[string]string),
		Validator: validator,
	}
}

//Получить список слов
func (dictionary *Dictionary) WordMap() map[string]string {
	return dictionary.wordMap
}

//Добавляет новое слово и его значение  в словарь
func (dictionary *Dictionary) AddWord(word, translate string) error {
	err := dictionary.Validator.Validate(word)
	if err != nil {
		return err
	}

	for value := range dictionary.wordMap {
		if value == word {
			return errors.New("такое слово уже есть в словаре")
		}
	}

	dictionary.wordMap[word] = translate
	return nil
}

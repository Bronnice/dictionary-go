package dictionary

import (
	"dictionary-go/validator"
	"errors"
)

// Словарь
type Dictionary struct {
	name      string
	wordMap   map[string]string
	validator validator.Validator
}

// Конструктор Dictionary
func NewDictionary(name string, validator validator.Validator) *Dictionary {
	return &Dictionary{
		name:      name,
		wordMap:   make(map[string]string),
		validator: validator,
	}
}

func (dictionary *Dictionary) Name() string {
	return dictionary.name
}

// Получить список слов
func (dictionary *Dictionary) WordMap() map[string]string {
	return dictionary.wordMap
}

// Добавляет новое слово и его значение в словарь
func (dictionary *Dictionary) AddWord(word, translate string) error {
	err := dictionary.validator.ValidateWord(word)
	if err != nil {
		return err
	}

	err = validator.ValidateTranslate(translate)
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

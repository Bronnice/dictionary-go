package validator

import (
	"errors"
	"regexp"
)

// Реализация интерфейса Validator, только латинские буквы
type EnglishWordValidator struct {
	wordCount int
}

// Конструктор для EnglishDictionaryValidator
func NewEnglishWordValidator() *EnglishWordValidator {
	return &EnglishWordValidator{
		wordCount: 4,
	}
}

// Проверяет подходит ли слово по критериям словаря (только английские символы, максимум 4 символа для слова)
func (validator EnglishWordValidator) ValidateWord(word string) error {
	if len(word) == 0 {
		return errors.New(EmptyWordMessage)
	}

	if len(word) > validator.wordCount {
		return errors.New(WordIsLongMessage)
	}

	isMatched, err := regexp.MatchString(`^[A-Za-z]+$`, word)
	if err != nil {
		return err
	}
	if isMatched {
		return nil
	}
	return errors.New(InvalidWordMessage)
}

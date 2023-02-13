package validator

import (
	"errors"
	"regexp"
)

// Реализация интерфейса Validator, только латинские буквы
type NumberWordValidator struct {
	wordCount int
}

// Конструктор NumberDictionaryValidator
func NewNumberWordValidator() *NumberWordValidator {
	return &NumberWordValidator{
		wordCount: 5,
	}
}

// Проверяет подходит ли слово по критериям словаря (только цифры, максимум 5 символа для слова)
func (validator *NumberWordValidator) ValidateWord(word string) error {
	if len(word) == 0 {
		return errors.New(EmptyWordMessage)
	}

	if len(word) > validator.wordCount {
		return errors.New(WordIsLongMessage)
	}

	isMatched, err := regexp.MatchString(`^[0-9]+$`, word)
	if err != nil {
		return err
	}
	if isMatched {
		return nil
	}
	return errors.New(InvalidWordMessage)
}

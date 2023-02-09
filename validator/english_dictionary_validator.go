package validator

import (
	"errors"
	"regexp"
)

// Реализация интерфейса Validator, только латинские буквы
type EnglishDictionaryValidator struct {
	wordCount int
}

// Конструктор для EnglishDictionaryValidator
func NewEnglishDictionaryValidator() *EnglishDictionaryValidator {
	return &EnglishDictionaryValidator{
		wordCount: 4,
	}
}

// Проверка подходит ли слово по критериям словаря (только английские символы, максимум 4 символа)
func (validator EnglishDictionaryValidator) ValidateWord(word string) error {
	if len(word) == 0 {
		return errors.New(EnterYourWordMessage)
	}

	if len(word) > validator.wordCount {
		return errors.New(WordIsLongMessage)
	}

	isMatched, err := regexp.MatchString(`[A-Za-z]`, word)
	if err != nil {
		return err
	}
	if isMatched {
		return nil
	}
	return errors.New(InvalidWordMessage)

}

package validator

import (
	"errors"
	"regexp"
	"strings"
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

// Проверяет подходит ли слово по критериям словаря (только английские символы, максимум 4 символа для слова)
func (validator EnglishDictionaryValidator) ValidateWord(word string) error {
	if len(strings.TrimSpace(word)) == 0 {
		return errors.New(EmptyInputMessage)
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

// Проверяет подходит ли перевод по критериям словаря (не только цифры, не менее 1 символа)
func (validator EnglishDictionaryValidator) ValidateTranslate(translate string) error {
	if len(strings.TrimSpace(translate)) < 1 {
		return errors.New(EmptyInputMessage)
	}

	isMatched, err := regexp.MatchString(`^[\d]+$`, translate)
	if err != nil {
		return err
	}
	if !isMatched {
		return nil
	}
	return errors.New(InvalidWordMessage)
}

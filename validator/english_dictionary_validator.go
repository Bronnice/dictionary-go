package validator

import (
	"dictionary-go/message"
	"errors"
	"regexp"
)

//Реализация интерфейса Validator, только латинские буквы
type EnglishDictionaryValidator struct {
	wordCount int
}

//Конструктор для EnglishDictionaryValidator
func NewEnglishDictionaryValidator() *EnglishDictionaryValidator {
	return &EnglishDictionaryValidator{
		wordCount: 4,
	}
}

//Проверка подходит ли слово по критериям словаря (только английские символы, максимум 4 символа)
func (validator EnglishDictionaryValidator) Validate(word string) error {
	if len(word) == 0 {
		return errors.New(message.EnterYourWordMessage)
	}

	if len(word) > validator.wordCount {
		return errors.New(message.WordIsLongMessage)
	}

	isMatched, err := regexp.MatchString(`[A-Za-z]`, word)
	if err != nil {
		return err
	}
	if isMatched {
		return nil
	} else {
		return errors.New(message.InvalidWordMessage + "Только буквы латинского алфавита!")
	}

}

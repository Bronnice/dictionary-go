package validator

import (
	"errors"
	"regexp"
)

//Реализация интерфейса Validator, только цифры
type NumbersDictionaryValidator struct {
	wordCount int
}

//Конструктор для NumbersDictionaryValidator
func NewNumbersDictionaryValidator() *NumbersDictionaryValidator {
	return &NumbersDictionaryValidator{
		wordCount: 5,
	}
}

//Проверка подходит ли слово по критериям словаря (только цифры, максимум 5 символов)
func (validator NumbersDictionaryValidator) Validate(word string) error {
	if len(word) == 0 {
		return errors.New(enterYourWordMessage)
	}

	if len(word) > validator.wordCount {
		return errors.New(wordIsLongMessage)
	}

	isMatched, err := regexp.MatchString(`\d`, word)
	if err != nil {
		return err
	}
	if isMatched {
		return nil
	} else {
		return errors.New(invalidWordMessage + "Только буквы цифры!")
	}

}

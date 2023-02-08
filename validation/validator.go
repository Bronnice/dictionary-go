package validation

import (
	"dictionary-go/enum"
	"dictionary-go/utils"
	"errors"
	"regexp"
)

var invalidWordMessage string = "Слово не подходит по критерию: "

type Validator struct {
	wordCount int
	wordType  enum.WordType
}

//Конструктор для Validator`a
func NewValidator(wordCount int, wordType enum.WordType) *Validator {
	return &Validator{
		wordCount: wordCount,
		wordType:  wordType,
	}
}

//Проверка подходит ли слово по критериям словаря
func IsWordValid(word *string, validator Validator) (*bool, error) {
	if len(*word) == 0 {
		return nil, errors.New("введите слово")
	}

	if len(*word) > validator.wordCount {
		return nil, errors.New("слово длиннее указанного размера")
	}

	switch validator.wordType {
	case "ENGLISH":
		matched, err := regexp.MatchString(`\D`, *word)
		if err != nil {
			return nil, err
		}
		if matched {
			return utils.GetPointer(true), nil
		} else {
			return nil, errors.New(invalidWordMessage + "Только буквы латинского алфавита")
		}
	case "NUMBERS":
		matched, err := regexp.MatchString(`\d`, *word)
		if err != nil {
			return nil, err
		}
		if matched {
			return utils.GetPointer(true), nil
		} else {
			return nil, errors.New(invalidWordMessage + "Только цифры")
		}
	default:
		return utils.GetPointer(false), nil
	}
}

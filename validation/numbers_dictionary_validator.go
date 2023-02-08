package validation

import (
	"dictionary-go/enum"
	"dictionary-go/utils"
	"errors"
	"regexp"
)

type NumbersDictionaryValidator struct {
	wordCount int
	wordType  enum.WordType
}

func NewNumbersDictionaryValidator() *NumbersDictionaryValidator {
	return &NumbersDictionaryValidator{
		wordCount: 5,
		wordType:  enum.NUMBERS,
	}
}

//Проверка подходит ли слово по критериям словаря (только цифры, максимум 5 символов)
func (validator NumbersDictionaryValidator) IsWordValid(word *string) (*bool, error) {
	if len(*word) == 0 {
		return nil, errors.New("введите слово")
	}

	if len(*word) > validator.wordCount {
		return nil, errors.New("слово длиннее указанного размера")
	}

	matched, err := regexp.MatchString(`\d`, *word)
	if err != nil {
		return nil, err
	}
	if matched {
		return utils.GetPointer(true), nil
	} else {
		return nil, errors.New("cлово не подходит по критерию: Только буквы цифры")
	}

}

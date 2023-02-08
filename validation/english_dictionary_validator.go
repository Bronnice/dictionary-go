package validation

import (
	"dictionary-go/enum"
	"dictionary-go/utils"
	"errors"
	"regexp"
)

type EnglishDictionaryValidator struct {
	wordCount int
	wordType  enum.WordType
}

func NewEnglishDictionaryValidator() *EnglishDictionaryValidator {
	return &EnglishDictionaryValidator{
		wordCount: 4,
		wordType:  enum.ENGLISH,
	}
}

//Проверка подходит ли слово по критериям словаря (только английские символы, максимум 4 символа)
func (validator EnglishDictionaryValidator) IsWordValid(word *string) (*bool, error) {
	if len(*word) == 0 {
		return nil, errors.New("введите слово")
	}

	if len(*word) > validator.wordCount {
		return nil, errors.New("слово длиннее указанного размера")
	}

	matched, err := regexp.MatchString(`[A-Za-z]`, *word)
	if err != nil {
		return nil, err
	}
	if matched {
		return utils.GetPointer(true), nil
	} else {
		return nil, errors.New("cлово не подходит по критерию: Только буквы латинского алфавита")
	}

}

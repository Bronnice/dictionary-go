package validator

import (
	"errors"
	"regexp"
)

// Проверяет подходит ли перевод по критериям словаря (не только цифры, не менее 1 символа)
func ValidateTranslate(translate string) error {
	if len(translate) < 1 {
		return errors.New(EmptyTranslateMessage)
	}

	isMatched, err := regexp.MatchString(`^[\d]+$`, translate)
	if err != nil {
		return err
	}
	if !isMatched {
		return nil
	}
	return errors.New(InvalidTranslateMessage)
}

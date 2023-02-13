package validator

import (
	"dictionary-go/validator"
	"testing"

	testSuite "github.com/stretchr/testify/suite"
)

// Тесты для метода ValidateTranslate EnglishWordValidator`a
type EnglishValidateWordTestSuite struct {
	testSuite.Suite
}

func Test_EnglishIsWordValidTestSuite(t *testing.T) {
	testSuite.Run(t, new(EnglishValidateWordTestSuite))
}

func (testSuite *EnglishValidateWordTestSuite) Test_ValidateWord_withValidWord_expectNoError() {
	word := "qwer"
	err := validator.NewEnglishWordValidator().ValidateWord(word)

	testSuite.NoError(err)
}

func (testSuite *EnglishValidateWordTestSuite) Test_ValidateWord_withInvalidWord_expectError() {
	words := []string{"1234", "qwert", "qwe1", "    ", "qwer123"}
	var err error
	validator := validator.NewEnglishWordValidator()

	for _, word := range words {
		err = validator.ValidateWord(word)

		testSuite.Error(err)
	}

}

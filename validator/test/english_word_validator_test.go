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
	wordErr := validator.NewEnglishWordValidator().ValidateWord(word)

	testSuite.NoError(wordErr)
}

func (testSuite *EnglishValidateWordTestSuite) Test_ValidateWord_withInvalidWord_expectError() {
	values := []string{"1234", "qwert"}

	for i := range values {
		word := values[i]
		err := validator.NewEnglishWordValidator().ValidateWord(word)

		testSuite.Error(err)
	}

}

package validator

import (
	"dictionary-go/validator"
	"testing"

	testSuite "github.com/stretchr/testify/suite"
)

// Тесты для метода ValidateWord EnglishWordValidator`a
type EnglishValidateWordTestSuite struct {
	testSuite.Suite
	validator validator.EnglishWordValidator
}

func Test_EnglishIsWordValidTestSuite(t *testing.T) {
	testSuite.Run(t, new(EnglishValidateWordTestSuite))
}

func (testSuite *EnglishValidateWordTestSuite) SetupSuite() {
	testSuite.validator = *validator.NewEnglishWordValidator()
}

func (testSuite *EnglishValidateWordTestSuite) Test_EnglishValidateWord_withValidWord_expectNoError() {
	word := "qwer"
	err := testSuite.validator.ValidateWord(word)

	testSuite.NoError(err)
}

func (testSuite *EnglishValidateWordTestSuite) Test_EnglishValidateWord_withInvalidWord_expectError() {
	words := []string{"1234", "qwert", "qwe1", "    ", "qwer123"}
	var err error

	for _, word := range words {
		err = testSuite.validator.ValidateWord(word)

		testSuite.Error(err)
	}
}

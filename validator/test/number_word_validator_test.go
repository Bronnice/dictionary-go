package validator

import (
	"dictionary-go/validator"
	"testing"

	testSuite "github.com/stretchr/testify/suite"
)

// Тесты для метода ValidateWord NumberWordValidator`a
type NumberValidateWordTestSuite struct {
	testSuite.Suite
	validator validator.NumberWordValidator
}

func Test_NumberIsWordValidTestSuite(t *testing.T) {
	testSuite.Run(t, new(NumberValidateWordTestSuite))
}

func (testSuite *NumberValidateWordTestSuite) SetupSuite() {
	testSuite.validator = *validator.NewNumberWordValidator()
}

func (testSuite *NumberValidateWordTestSuite) Test_NumberValidateWord_withValidWord_expectNoError() {
	word := "12345"
	err := testSuite.validator.ValidateWord(word)

	testSuite.NoError(err)
}

func (testSuite *NumberValidateWordTestSuite) Test_NumberValidateWord_withInvalidWord_expectError() {
	words := []string{"awd", "123456", "qwe12", "    ", "qwer123"}
	var err error

	for _, word := range words {
		err = testSuite.validator.ValidateWord(word)

		testSuite.Error(err)
	}
}

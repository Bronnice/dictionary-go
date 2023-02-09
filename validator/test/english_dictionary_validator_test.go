package validator

import (
	"dictionary-go/validator"
	"testing"

	testSuite "github.com/stretchr/testify/suite"
)

//Тесты для метода IsWordValid EnglishDictionaryValidator`a
type EnglishDictionaryIsWordValidTestSuite struct {
	testSuite.Suite
}

func Test_EnglishDictionaryIsWordValidTestSuite(t *testing.T) {
	testSuite.Run(t, new(EnglishDictionaryIsWordValidTestSuite))
}

func (testSuite *EnglishDictionaryIsWordValidTestSuite) Test_IsWordValid_withValidWord_expectNoError() {
	word := "qwer"
	err := validator.NewEnglishDictionaryValidator().ValidateWord(word)

	testSuite.NoError(err)
}

func (testSuite *EnglishDictionaryIsWordValidTestSuite) Test_IsWordValid_withInvalidWord_expectError() {
	values := []string{"1234", "qwert"}

	for i := range values {
		word := values[i]
		err := validator.NewEnglishDictionaryValidator().ValidateWord(word)

		testSuite.NotNil(err)
	}

}

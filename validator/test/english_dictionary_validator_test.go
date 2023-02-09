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

func (testSuite *EnglishDictionaryIsWordValidTestSuite) Test_IsWordValid_withInvalidWord_expectError() {
	values := [2]string{"1234", "qwert"}

	for i := 0; i < len(values); i++ {
		word := values[i]
		err := validator.NewEnglishDictionaryValidator().Validate(word)

		testSuite.NotNil(err)
	}

}

func (testSuite *EnglishDictionaryIsWordValidTestSuite) Test_IsWordValid_withValidWord_expectTrue() {
	word := "qwer"
	err := validator.NewEnglishDictionaryValidator().Validate(word)

	testSuite.Nil(err)
}

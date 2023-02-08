package tests

import (
	"dictionary-go/validation"
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

func (testSuite *EnglishDictionaryIsWordValidTestSuite) Test_IsWordValid_withInvalidWord_expectedError() {
	word := "1234"
	result, err := validation.NewEnglishDictionaryValidator().IsWordValid(&word)

	testSuite.Nil(result)
	testSuite.NotNil(err)
}

func (testSuite *EnglishDictionaryIsWordValidTestSuite) Test_IsWordValid_withValidWord_expectedTrue() {
	word := "qwer"
	result, err := validation.NewEnglishDictionaryValidator().IsWordValid(&word)

	testSuite.NotNil(result)
	testSuite.Nil(err)
}

func (testSuite *EnglishDictionaryIsWordValidTestSuite) Test_IsWordValid_withLongLenght_expectedError() {
	word := "qwert"
	result, err := validation.NewEnglishDictionaryValidator().IsWordValid(&word)

	testSuite.Nil(result)
	testSuite.NotNil(err)
}

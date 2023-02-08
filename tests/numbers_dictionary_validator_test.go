package tests

import (
	"dictionary-go/validation"
	"testing"

	testSuite "github.com/stretchr/testify/suite"
)

//Тесты для метода IsWordValid NumbersDictionaryValidator`a
type NumbersDictionaryIsWordValidTestSuite struct {
	testSuite.Suite
}

func Test_NumbersDictionaryIsWordValidTestSuite(t *testing.T) {
	testSuite.Run(t, new(EnglishDictionaryIsWordValidTestSuite))
}

func (testSuite *NumbersDictionaryIsWordValidTestSuite) Test_IsWordValid_withInvalidWord_expectedError() {
	word := "qwe"
	result, err := validation.NewNumbersDictionaryValidator().IsWordValid(&word)

	testSuite.Nil(result)
	testSuite.NotNil(err)
}

func (testSuite *NumbersDictionaryIsWordValidTestSuite) Test_IsWordValid_withValidWord_expectedTrue() {
	word := "1234"
	result, err := validation.NewNumbersDictionaryValidator().IsWordValid(&word)

	testSuite.NotNil(result)
	testSuite.Nil(err)
}

func (testSuite *NumbersDictionaryIsWordValidTestSuite) Test_IsWordValid_withLongLenght_expectedError() {
	word := "1234"
	result, err := validation.NewNumbersDictionaryValidator().IsWordValid(&word)

	testSuite.Nil(result)
	testSuite.NotNil(err)
}

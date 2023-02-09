package test

import (
	"dictionary-go/validator"
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

func (testSuite *NumbersDictionaryIsWordValidTestSuite) Test_IsWordValid_withInvalidWord_expectError() {
	word := "qwe"
	err := validator.NewNumbersDictionaryValidator().Validate(word)

	testSuite.NotNil(err)

	word = "1234"
	err = validator.NewNumbersDictionaryValidator().Validate(word)

	testSuite.NotNil(err)
}

func (testSuite *NumbersDictionaryIsWordValidTestSuite) Test_IsWordValid_withValidWord_expectTrue() {
	word := "1234"
	err := validator.NewNumbersDictionaryValidator().Validate(word)

	testSuite.Nil(err)
}

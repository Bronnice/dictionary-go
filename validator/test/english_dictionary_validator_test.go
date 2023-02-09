package validator

import (
	"dictionary-go/validator"
	"testing"

	testSuite "github.com/stretchr/testify/suite"
)

// Тесты для метода IsWordValid EnglishDictionaryValidator`a
type EnglishDictionaryIsWordValidTestSuite struct {
	testSuite.Suite
}

func Test_EnglishDictionaryIsWordValidTestSuite(t *testing.T) {
	testSuite.Run(t, new(EnglishDictionaryIsWordValidTestSuite))
}

func (testSuite *EnglishDictionaryIsWordValidTestSuite) Test_IsWordValid_withValidWord_expectNoError() {
	word := "qwer"
	translate := "awdss"
	wordErr := validator.NewEnglishDictionaryValidator().ValidateWord(word)
	translateErr := validator.NewEnglishDictionaryValidator().ValidateTranslate(translate)

	testSuite.NoError(wordErr)
	testSuite.NoError(translateErr)
}

func (testSuite *EnglishDictionaryIsWordValidTestSuite) Test_IsWordValid_withInvalidWord_expectError() {
	wordValues := []string{"1234", "qwert"}
	translateValues := []string{"1232", ""}

	for i := range wordValues {
		word := wordValues[i]
		translate := translateValues[i]
		wordErr := validator.NewEnglishDictionaryValidator().ValidateWord(word)
		translateErr := validator.NewEnglishDictionaryValidator().ValidateTranslate(translate)

		testSuite.NoError(wordErr)
		testSuite.NoError(translateErr)
	}

}

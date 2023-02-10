package dictionary

import (
	"dictionary-go/dictionary"
	"dictionary-go/validator"
	"testing"

	testSuite "github.com/stretchr/testify/suite"
)

// Тесты для метода AddWord Dictionary
type AddWordTestSuite struct {
	testSuite.Suite
	dictionary dictionary.Dictionary
}

func Test_AddWordTestSuite(t *testing.T) {
	testSuite.Run(t, new(AddWordTestSuite))
}

func (testSuite *AddWordTestSuite) SetupTest() {
	validator := validator.NewEnglishWordValidator()
	testSuite.dictionary = *dictionary.NewDictionary(validator)
}

func (testSuite *AddWordTestSuite) Test_AddWord_withNewWord_expectNoError() {
	word := "word"
	translate := "translate"
	err := testSuite.dictionary.AddWord(word, translate)

	testSuite.Contains(testSuite.dictionary.WordMap(), word)
	testSuite.Contains(testSuite.dictionary.WordMap()[word], translate)

	testSuite.NoError(err)
}

func (testSuite *AddWordTestSuite) Test_AddWord_withAlreadyExistedWord_expectError() {
	word := "word"
	translate := "translate"
	var err error

	for i := 0; i < 2; i++ {
		err = testSuite.dictionary.AddWord(word, translate)
	}
	testSuite.Contains(testSuite.dictionary.WordMap(), word)
	testSuite.Contains(testSuite.dictionary.WordMap()[word], translate)
	testSuite.Error(err)
}

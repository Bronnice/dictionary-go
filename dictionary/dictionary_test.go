package dictionary

import (
	"dictionary-go/validator"
	"testing"

	testSuite "github.com/stretchr/testify/suite"
)

// Тесты для метода AddWord Dictionary
type AddWordTestSuite struct {
	testSuite.Suite
	dictionary Dictionary
	validator  validator.EnglishWordValidator
}

func Test_AddWordTestSuite(t *testing.T) {
	testSuite.Run(t, new(AddWordTestSuite))
}

func (testSuite *AddWordTestSuite) SetupSuite() {
	testSuite.validator = *validator.NewEnglishWordValidator()
}

func (testSuite *AddWordTestSuite) SetupTest() {
	testSuite.dictionary = *NewDictionary("testDictionary", &testSuite.validator)
}

func (testSuite *AddWordTestSuite) Test_AddWord_withNewWord_expectNoError() {
	word := "word"
	translate := "translate"
	err := testSuite.dictionary.AddWord(word, translate)

	testSuite.NoError(err)
	testSuite.Len(testSuite.dictionary.WordMap(), 1)
	testSuite.Contains(testSuite.dictionary.WordMap(), word)
	testSuite.Contains(testSuite.dictionary.WordMap()[word], translate)
}

func (testSuite *AddWordTestSuite) Test_AddWord_withAlreadyExistedWord_expectError() {
	word := "word"
	translate := "translate"
	var err error

	for i := 0; i < 2; i++ {
		err = testSuite.dictionary.AddWord(word, translate)
	}

	testSuite.Error(err)
	testSuite.Len(testSuite.dictionary.WordMap(), 1)
	testSuite.Contains(testSuite.dictionary.WordMap(), word)
	testSuite.Contains(testSuite.dictionary.WordMap()[word], translate)
}

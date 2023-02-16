package dictionary

import (
	"dictionary-go/validator"
	"testing"

	testSuite "github.com/stretchr/testify/suite"
)

// Тесты для метода AddWord Dictionary
type DictionaryTestSuite struct {
	testSuite.Suite
	dictionary Dictionary
	validator  validator.EnglishWordValidator
	word       string
	translate  string
}

func Test_DictionaryTestSuite(t *testing.T) {
	testSuite.Run(t, new(DictionaryTestSuite))
}

func (testSuite *DictionaryTestSuite) SetupSuite() {
	testSuite.validator = *validator.NewEnglishWordValidator()
	testSuite.dictionary = *NewDictionary("testDictionary", &testSuite.validator)

	testSuite.word = "word"
	testSuite.translate = "translate"
}

func (testSuite *DictionaryTestSuite) Test_AddWord_withANewWord_expectNoError() {
	err := testSuite.dictionary.AddWord(testSuite.word, testSuite.translate)

	testSuite.NoError(err)
	testSuite.Len(testSuite.dictionary.WordMap(), 1)
	testSuite.Contains(testSuite.dictionary.WordMap(), testSuite.word)
	testSuite.Contains(testSuite.dictionary.WordMap()[testSuite.word], testSuite.translate)
}

func (testSuite *DictionaryTestSuite) Test_AddWord_withAlreadyExistedWord_expectError() {
	err := testSuite.dictionary.AddWord(testSuite.word, testSuite.translate)

	testSuite.Error(err)
	testSuite.Len(testSuite.dictionary.WordMap(), 1)
	testSuite.Contains(testSuite.dictionary.WordMap(), testSuite.word)
	testSuite.Contains(testSuite.dictionary.WordMap()[testSuite.word], testSuite.translate)
}

func (testSuite *DictionaryTestSuite) Test_SearchPairByWord_withExistedWord_expectNoError() {
	pair, err := testSuite.dictionary.SearchPairByWord(testSuite.word)

	testSuite.NoError(err)
	testSuite.NotEmpty(pair)
}

func (testSuite *DictionaryTestSuite) Test_SearchPairByWord_withNonexistedWord_expectError() {
	name := "dorw"

	pair, err := testSuite.dictionary.SearchPairByWord(name)

	testSuite.Error(err)
	testSuite.Empty(pair)
}

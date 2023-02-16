package dictionary

import (
	"dictionary-go/validator"
	"testing"

	testSuite "github.com/stretchr/testify/suite"
)

// Тесты для методов Dictionary
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

	testSuite.word = "word"
	testSuite.translate = "translate"
}

func (testSuite *DictionaryTestSuite) SetupTest() {
	testSuite.dictionary = *NewDictionary("testDictionary", &testSuite.validator)
}

func (testSuite *DictionaryTestSuite) Test_AddWord_withNewWord_expectNoError() {
	err := testSuite.dictionary.AddWord(testSuite.word, testSuite.translate)

	testSuite.NoError(err)
	testSuite.Len(testSuite.dictionary.WordMap(), 1)
	testSuite.Contains(testSuite.dictionary.WordMap(), testSuite.word)
	testSuite.Contains(testSuite.dictionary.WordMap()[testSuite.word], testSuite.translate)
}

func (testSuite *DictionaryTestSuite) Test_AddWord_withAlreadyExistedWord_expectError() {
	var err error

	for i := 0; i < 2; i++ {
		err = testSuite.dictionary.AddWord(testSuite.word, testSuite.translate)
	}

	testSuite.Error(err)
	testSuite.Len(testSuite.dictionary.WordMap(), 1)
	testSuite.Contains(testSuite.dictionary.WordMap(), testSuite.word)
	testSuite.Contains(testSuite.dictionary.WordMap()[testSuite.word], testSuite.translate)
}

func (testSuite *DictionaryTestSuite) Test_SearchTranslateByWord_withExistedWord_expectNotNil() {
	testSuite.dictionary.AddWord(testSuite.word, testSuite.translate)
	desiredTranslate := testSuite.dictionary.SearchTranslateByWord(testSuite.word)

	testSuite.NotNil(desiredTranslate)
}

func (testSuite *DictionaryTestSuite) Test_SearchTranslateByWord_withNonexistedWord_expectNil() {
	name := "dorw"

	testSuite.dictionary.AddWord(testSuite.word, testSuite.translate)
	desiredTranslate := testSuite.dictionary.SearchTranslateByWord(name)

	testSuite.Nil(desiredTranslate)
}

func (testSuite *DictionaryTestSuite) Test_DeletePairByWord_withExistedWord_expectNoError() {
	testSuite.dictionary.AddWord(testSuite.word, testSuite.translate)

	err := testSuite.dictionary.DeletePairByWord(testSuite.word)
	testSuite.NoError(err)
	testSuite.Empty(testSuite.dictionary.wordMap)
}

func (testSuite *DictionaryTestSuite) Test_DeletePairByWord_withNonexistedWord_expectError() {
	testSuite.dictionary.AddWord(testSuite.word, testSuite.translate)

	err := testSuite.dictionary.DeletePairByWord("dorw")
	testSuite.Error(err)
	testSuite.NotEmpty(testSuite.dictionary.wordMap)
}

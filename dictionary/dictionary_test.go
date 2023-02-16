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

	testSuite.word = "word"
	testSuite.translate = "translate"
}

func (testSuite *DictionaryTestSuite) SetupTest() {
	testSuite.dictionary = *NewDictionary("testDictionary", &testSuite.validator)
}

func (testSuite *DictionaryTestSuite) Test_AddWord_withANewWord_expectNoError() {
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

func (testSuite *DictionaryTestSuite) Test_SearchPairByWord_withExistedWord_expectNoError() {
	err := testSuite.dictionary.AddWord(testSuite.word, testSuite.translate)
	desiredTranslate := testSuite.dictionary.SearchPairByWord(testSuite.word)

	testSuite.NoError(err)
	testSuite.NotEmpty(desiredTranslate)
}

func (testSuite *DictionaryTestSuite) Test_SearchPairByWord_withNonexistedWord_expectError() {
	name := "dorw"

	err := testSuite.dictionary.AddWord(testSuite.word, testSuite.translate)
	desiredTranslate := testSuite.dictionary.SearchPairByWord(name)

	testSuite.NoError(err)
	testSuite.Empty(desiredTranslate)
}

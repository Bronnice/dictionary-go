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
}

func Test_AddWordTestSuite(t *testing.T) {
	testSuite.Run(t, new(AddWordTestSuite))
}

func (testSuite *AddWordTestSuite) Test_AddWord_withNewWord_expectNoError() {
	dictionary := dictionary.NewDictionary(validator.NewEnglishWordValidator())

	err := dictionary.AddWord("word", "translate")

	testSuite.NoError(err)
}

func (testSuite *AddWordTestSuite) Test_AddWord_withAlreadyExistedWord_expectError() {
	dictionary := dictionary.NewDictionary(validator.NewEnglishWordValidator())
	word := "word"
	var err error

	for i := 0; i < 2; i++ {
		err = dictionary.AddWord(word, "translate")
	}

	testSuite.Error(err)
}

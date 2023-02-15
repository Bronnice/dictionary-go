package datasource

import (
	"dictionary-go/dictionary"
	"testing"

	testSuite "github.com/stretchr/testify/suite"
)

// Тесты для метода GetDictionaryByName Datasource
type GetDictionaryByNameTestSuite struct {
	testSuite.Suite
}

func Test_GetDictionaryByNameTestSuite(t *testing.T) {
	testSuite.Run(t, new(GetDictionaryByNameTestSuite))
}

func (testSuite *GetDictionaryByNameTestSuite) Test_GetDictionaryByName_withExistedName_expectEqual() {
	existedNames := []string{"Numbers", "English", "numbers", "english", "NUMBERS", "ENGLISH", "nuMbers", "enGlish"}
	var selectedDictionary *dictionary.Dictionary

	for _, name := range existedNames {
		selectedDictionary = GetDictionaryByName(name)

		testSuite.NotNil(selectedDictionary)
	}
}

func (testSuite *GetDictionaryByNameTestSuite) Test_GetDictionaryByName_withNotExistedName_expectNotEqual() {
	existedNames := []string{"awder", "Awdre", "aWdqe"}
	var selectedDictionary *dictionary.Dictionary

	for _, name := range existedNames {
		selectedDictionary = GetDictionaryByName(name)

		testSuite.Nil(selectedDictionary)
	}
}

package datasource

import (
	"dictionary-go/dictionary"
	"testing"

	testSuite "github.com/stretchr/testify/suite"
)

// Тесты для методов Datasource
type DataSourceTestSuite struct {
	testSuite.Suite
}

func Test_GetDictionaryByNameTestSuite(t *testing.T) {
	testSuite.Run(t, new(DataSourceTestSuite))
}

func (testSuite *DataSourceTestSuite) Test_GetDictionaryByName_withExistedName_expectEqual() {
	existedNames := []string{"Numbers", "English", "numbers", "english", "NUMBERS", "ENGLISH", "nuMbers", "enGlish"}
	var selectedDictionary *dictionary.Dictionary

	for _, name := range existedNames {
		selectedDictionary = GetDictionaryByName(name)

		testSuite.NotNil(selectedDictionary)
	}
}

func (testSuite *DataSourceTestSuite) Test_GetDictionaryByName_withNotExistedName_expectNotEqual() {
	nonExistedNames := []string{"awder", "Awdre", "aWdqe"}
	var selectedDictionary *dictionary.Dictionary

	for _, name := range nonExistedNames {
		selectedDictionary = GetDictionaryByName(name)

		testSuite.Nil(selectedDictionary)
	}
}

func (testSuite *DataSourceTestSuite) Test_GetDictionaryNames_expectEqualAndNotEmpty() {
	dictionaryNames := GetDictionariesNames()

	testSuite.NotEmpty(dictionaryNames)
	for i, dictionary := range dictionaries {
		testSuite.Equal(dictionaryNames[i], dictionary.Name())
	}
}

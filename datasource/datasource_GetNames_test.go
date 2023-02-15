package datasource

import (
	"testing"

	testSuite "github.com/stretchr/testify/suite"
)

// Тесты для метода AddWord Dictionary
type GetDictionaryNamesTestSuite struct {
	testSuite.Suite
}

func Test_GetDictionaryNamesTestSuite(t *testing.T) {
	testSuite.Run(t, new(GetDictionaryNamesTestSuite))
}

func (testSuite *GetDictionaryNamesTestSuite) Test_GetDictionaryNames() {
	dictionaryNames := GetDictionariesNames()

	for i, dictionary := range dictionaries {
		testSuite.Equal(dictionaryNames[i], dictionary.Name())
	}
}

package validator

import (
	"dictionary-go/validator"
	"testing"

	testSuite "github.com/stretchr/testify/suite"
)

// Тесты для метода ValidateTranslate
type ValidateTranslateTestSuite struct {
	testSuite.Suite
}

func Test_ValidateTranslateTestSuite(t *testing.T) {
	testSuite.Run(t, new(ValidateTranslateTestSuite))
}

func (testSuite *ValidateTranslateTestSuite) Test_ValidateTranslate_withValidTranslate_expectNoError() {
	translates := []string{"example", "example with numbers 123"}
	var err error

	for _, translate := range translates {
		err = validator.ValidateTranslate(translate)

		testSuite.NoError(err)
	}

}

func (testSuite *ValidateTranslateTestSuite) Test_ValidateTranslate_withInvalidTranslate_expectError() {
	translates := []string{"1232", ""}
	var err error

	for _, translate := range translates {
		err = validator.ValidateTranslate(translate)

		testSuite.Error(err)
	}

}

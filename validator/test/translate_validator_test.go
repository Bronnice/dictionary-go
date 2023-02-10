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
	values := []string{"example", "example with numbers 123"}

	for _, value := range values {
		translate := value
		err := validator.ValidateTranslate(translate)

		testSuite.NoError(err)
	}

}

func (testSuite *ValidateTranslateTestSuite) Test_ValidateTranslate_withInvalidTranslate_expectError() {
	values := []string{"1232", ""}

	for i := range values {
		translate := values[i]
		err := validator.ValidateTranslate(translate)

		testSuite.Error(err)
	}

}

package test

import (
	"dictionary-go/ui/console"
	"testing"

	testSuite "github.com/stretchr/testify/suite"
)

//Тесты для метода formatInput
type formatInputTestSuite struct {
	testSuite.Suite
}

func Test_formatInputTestSuite(t *testing.T) {
	testSuite.Run(t, new(formatInputTestSuite))
}

func (testSuite *formatInputTestSuite) Test_formatInput_withUnformated_expectFormated() {
	input := " ExAmPle "

	testSuite.Equal("example", console.FormatInput(input))
}

func (testSuite *formatInputTestSuite) Test_formatInput_withWhitespaces_expectEmptyString() {
	input := "    "

	testSuite.Equal("", console.FormatInput(input))
}

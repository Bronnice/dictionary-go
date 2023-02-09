package validator

type Validator interface {
	Validate(string) error
}

var enterYourWordMessage string = "Введите слово!"
var invalidWordMessage string = "Cлово не подходит по критерию: "
var wordIsLongMessage string = "Слово длиннее указанного размера!"

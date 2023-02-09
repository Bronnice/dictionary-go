package validator

//Валидация словаря
type Validator interface {
	Validate(string) error
}

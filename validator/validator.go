package validator

type Validator interface {
	Validate(string) error
}

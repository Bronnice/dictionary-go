package validation

type Validator interface {
	IsWordValid(*string) (*bool, error)
}

package validator

// Валидатор словаря
type Validator interface {
	// Валидация слова
	ValidateWord(string) error
}

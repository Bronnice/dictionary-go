package validator

// Валидатор словаря
type Validator interface {
	// Валидация слова
	ValidateWord(word string) error
}

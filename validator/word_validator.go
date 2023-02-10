package validator

// Валидатор словаря
type Validator interface {
	// Валидация вводимой слова
	ValidateWord(word string) error
}

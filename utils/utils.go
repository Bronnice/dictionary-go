package utils

//Возвращает указатель значения
func GetPointer[T any](value T) *T {
	return &value
}

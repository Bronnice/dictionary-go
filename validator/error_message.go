package validator

//Сообщения об ошибках ввода слова
var (
	EmptyWordMessage   = "Слово не может быть пустым!"
	InvalidWordMessage = "Слово не подходит по критерию словаря!"
	WordIsLongMessage  = "Слово длиннее, указанного в словаре размера!"
)

//Сообщения об ошибках ввода перевода
var (
	EmptyTranslateMessage   = "Перевод не может быть пустым!"
	InvalidTranslateMessage = "Перевод не может состоять только из цифр!"
)

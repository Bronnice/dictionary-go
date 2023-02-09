package console

import (
	"dictionary-go/dictionary"
	"dictionary-go/validator"
	"strings"
)

var helpMessage string = "Введите help - для помощи"

// Запуск UI
func RunUi() {
	validator := validator.NewEnglishDictionaryValidator()
	dictionary := dictionary.NewDictionary(validator)

	Println(helpMessage)
	for {
		input, err := ReadLine()
		if err != nil {
			Println(err.Error())
			return
		}

		input = formatInput(input)

		switch input {
		case "exit":
			return
		case "help":
			help()
		case "print":
			PrintDictionary(dictionary)
		case "switch":
			Println("Work in progress")
		case "add":
			add(dictionary)
		default:
			Println("Неизвестная команда, " + helpMessage)
		}
	}
}

// Форматирование пользовательского ввода
func formatInput(input string) string {
	return strings.TrimSpace(strings.ToLower(input))
}

// Отображение списка команд в коносль
func help() {
	Println("print  - просмотр словаря")
	Println("list - список всех словарей")
	Println("exit - прекратить работу")
	Println("switch - переключить словарь")
	Println("add - добавить новую пару слово - перевод в словарь")
}

// Добавление новой пары в словарь
func add(dictionary *dictionary.Dictionary) {
	Println("Введите слово: ")
	word, err := ReadLine()
	if err != nil {
		Println(err.Error())
		return
	}

	Println("Введите перевод: ")
	translate, err := ReadLine()
	if err != nil {
		Println(err.Error())
		return
	}

	err = dictionary.AddWord(word, translate)
	if err != nil {
		Println(err.Error())
		return
	}

	Println("Запись добавлена!")
}

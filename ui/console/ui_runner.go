package console

import (
	"dictionary-go/dictionary"
	"strings"
)

var helpMessage string = "Введите help - для помощи"
var inputIssueMessage string = "Ошибка ввода!"

//Запуск UI
func RunUi() {
	dictionary := dictionary.NewDictionary()

	Println(helpMessage)
	for {
		input, err := ReadLine()
		if err != nil {
			Println("Ошибка ввода!")
			return
		}

		input = FormatInput(input)

		switch input {
		case "exit":
			return
		case "help":
			help()
		case "print":
			PrintDictionary(dictionary)
		case "select":
			Println("Work in progress")
		case "add":
			add(*dictionary)
		default:
			Println("Неизвестная команда, " + helpMessage)
		}
	}
}

//Форматирование пользовательского ввода
func FormatInput(input string) string {
	return strings.TrimSpace(strings.ToLower(input))
}

//Отображение списка команд в коносль
func help() {
	Println("print  - просмотр словаря")
	Println("list - список всех словарей")
	Println("exit - прекратить работу")
	Println("select - переключить словарь")
}

//Добавление новой пары в словарь
func add(dictionary dictionary.Dictionary) {
	Println("Введите слово: ")
	word, err := ReadLine()
	if err != nil {
		Println(inputIssueMessage)
		return
	}
	Println("Введите перевод: ")
	translate, err := ReadLine()
	if err != nil {
		Println(inputIssueMessage)
		return
	}
	dictionary.AddWord(word, translate)
	Println("Запись добавлена!")
}

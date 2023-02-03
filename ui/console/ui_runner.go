package console

import (
	"dictionary-go/dictionary"
	"strings"
)

var printHelp string = "Введите help - для помощи"

//Запуск UI
func RunUi() {
	dictionary := dictionary.NewDictionary()

	Println(printHelp)
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
		default:
			Println("Неизвестная команда, " + printHelp)
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

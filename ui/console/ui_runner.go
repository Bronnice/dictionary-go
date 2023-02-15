package console

import (
	"dictionary-go/datasource"
	"dictionary-go/dictionary"
	"strings"
)

var selectedDictionary *dictionary.Dictionary = nil

// Запуск UI
func RunUi() {
	selectDictionary()
	help()
	for {
		input, err := ReadLine()
		if err != nil {
			Println(err.Error())
			return
		}

		switch strings.ToLower(input) {
		case "exit":
			return
		case "help":
			help()
		case "print":
			if selectedDictionary == nil {
				Println(chooseDictionaryMessage)
				continue
			}
			PrintDictionary(selectedDictionary)
		case "select":
			selectDictionary()
		case "add":
			if selectedDictionary == nil {
				Println(chooseDictionaryMessage)
				continue
			}
			add(selectedDictionary)
		default:
			Println("Неизвестная команда, " + helpMessage)
		}
	}
}

// Отображение списка команд в коносль
func help() {
	Println("help  - список доступных команд")
	Println("print  - просмотр словаря")
	Println("exit - прекратить работу")
	Println("select - выбрать словарь")
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

func selectDictionary() {
	for {
		Println("\nВыберите словарь(введите его имя):")
		for _, dictionary := range *datasource.Dictionaries() {
			Println(dictionary.Name())
		}

		input, err := ReadLine()
		if err != nil {
			Println(err.Error())
		}
		for _, dictionary := range *datasource.Dictionaries() {
			if strings.EqualFold(input, dictionary.Name()) {
				*selectedDictionary = dictionary
				Println(chosenDictionaryMessage + dictionary.Name())
			}
		}
		Println("Cловарь не существует!")
	}
}

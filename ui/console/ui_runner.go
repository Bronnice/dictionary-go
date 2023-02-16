package console

import (
	"dictionary-go/datasource"
	"dictionary-go/dictionary"
	"strings"
)

var selectedDictionary *dictionary.Dictionary = nil

// Запуск UI
func RunUi() {
	err := selectDictionary()
	if err != nil {
		Println(err.Error())
	}

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
			PrintDictionary(selectedDictionary)
		case "select":
			err := selectDictionary()
			if err != nil {
				Println(err.Error())
			}
		case "add":
			add(selectedDictionary)
		default:
			PrintlnFormatted("Неизвестная команда, %s", helpMessage)
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

func selectDictionary() error {
	for {
		Println("Выберите словарь(введите его имя):")
		for _, dictionaryName := range datasource.GetDictionariesNames() {
			Println(dictionaryName)
		}

		input, err := ReadLine()
		if err != nil {
			Println(err.Error())
			continue
		}

		if datasource.GetDictionaryByName(input) != nil {
			selectedDictionary = datasource.GetDictionaryByName(input)
			PrintlnFormatted(chosenDictionaryMessagePattern, selectedDictionary.Name())
			return nil
		}

		Println("Словарь не существует!")
	}
}

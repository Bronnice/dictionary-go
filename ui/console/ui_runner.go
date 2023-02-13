package console

import (
	"dictionary-go/dictionary"
	"dictionary-go/validator"
	"strings"
)

var dictionaries []dictionary.Dictionary = []dictionary.Dictionary{
	*dictionary.NewDictionary("English", validator.NewEnglishWordValidator()),
	*dictionary.NewDictionary("Numbers", validator.NewNumberWordValidator()),
}
var chosenDictionary int = -1

// Запуск UI
func RunUi() {
	Println(helpMessage)
	switchCommand()
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
			if chosenDictionary == -1 {
				Println(chooseDictionaryMessage)
				continue
			}
			PrintDictionary(&dictionaries[chosenDictionary])
		case "switch":
			switchCommand()
		case "add":
			if chosenDictionary == -1 {
				Println(chooseDictionaryMessage)
				continue
			}
			add(&dictionaries[chosenDictionary])
		default:
			Println("Неизвестная команда, " + helpMessage)
		}
	}
}

// Отображение списка команд в коносль
func help() {
	Println("print  - просмотр словаря")
	Println("exit - прекратить работу")
	Println("switch - выбрать словарь")
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

func switchCommand() {
	Println("Выберите словарь(введите его имя):")
	for i := range dictionaries {
		Println(dictionaries[i].Name())
	}

	input, err := ReadLine()
	if err != nil {
		Println(err.Error())
	}

	switch strings.ToLower(input) {
	case "english":
		chosenDictionary = 0
		Println(chosenDictionaryMessage)
	case "numbers":
		chosenDictionary = 1
		Println(chosenDictionaryMessage)
	default:
		chosenDictionary = -1
		Println("Такого словаря не существует!")
	}
}

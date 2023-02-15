package datasource

import (
	"dictionary-go/dictionary"
	"dictionary-go/validator"
	"strings"
)

var dictionaries []dictionary.Dictionary = []dictionary.Dictionary{
	*dictionary.NewDictionary("English", validator.NewEnglishWordValidator()),
	*dictionary.NewDictionary("Numbers", validator.NewNumberWordValidator()),
}

func Dictionaries() *[]dictionary.Dictionary {
	return &dictionaries
}

func SelectDictionary(name string) *dictionary.Dictionary {
	for _, dictionary := range dictionaries {
		if strings.EqualFold(name, dictionary.Name()) {
			return &dictionary
		}
	}
	return nil
}

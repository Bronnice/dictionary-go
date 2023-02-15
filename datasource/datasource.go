package datasource

import (
	"dictionary-go/dictionary"
	"dictionary-go/validator"
)

var dictionaries []dictionary.Dictionary = []dictionary.Dictionary{
	*dictionary.NewDictionary("English", validator.NewEnglishWordValidator()),
	*dictionary.NewDictionary("Numbers", validator.NewNumberWordValidator()),
}

func Dictionaries() *[]dictionary.Dictionary {
	return &dictionaries
}

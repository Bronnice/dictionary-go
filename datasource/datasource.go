package datasource

import (
	"dictionary-go/dictionary"
)

func AddNewIssue(dictionary dictionary.Dictionary, word, translate string) {
	dictionary.WordMap()[word] = translate
}

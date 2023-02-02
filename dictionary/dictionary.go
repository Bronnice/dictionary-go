package dictionary

type Dictionary struct {
	name       string
	dictionary map[string]string
}

func CreateDictionary(dictionaryName string) Dictionary {
	return Dictionary{
		name:       dictionaryName,
		dictionary: make(map[string]string),
	}
}

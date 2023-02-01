package dictionary

//Структура для словаря
type Dictionary struct {
	Name    string
	wordMap map[string]string
}

//Конструктор для Dictionary
func NewDictionary(name string) *Dictionary {
	return &Dictionary{
		Name:    name,
		wordMap: make(map[string]string),
	}
}

//Получить список слов
func (dic *Dictionary) WordMap() map[string]string {
	return dic.wordMap
}

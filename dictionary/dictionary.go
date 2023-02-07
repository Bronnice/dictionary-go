package dictionary

//Структура для словаря
type Dictionary struct {
	Number  int
	wordMap map[string]string
}

//Конструктор для Dictionary
func NewDictionary() *Dictionary {
	return &Dictionary{
		Number:  -1,
		wordMap: make(map[string]string),
	}
}

//Получить список слов
func (dic *Dictionary) WordMap() map[string]string {
	return dic.wordMap
}

//Добавляет новое слово и его значение  в словарь
func (dictionary Dictionary) AddWord(word, translate string) {
	dictionary.WordMap()[word] = translate
}

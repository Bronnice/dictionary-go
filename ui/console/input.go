package console

import (
	"bufio"
	"os"
	"strings"
)

// Чтение строки из консоли
func ReadLine() (string, error) {
	in := bufio.NewReader(os.Stdin)
	str, err := in.ReadString('\n')
	return formatInput(str), err
}

// Форматирование пользовательского ввода
func formatInput(input string) string {
	return strings.TrimSpace(strings.Trim(input, "\n"))
}

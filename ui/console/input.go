package console

import (
	"bufio"
	"os"
	"strings"
)

//Чтение строки из консоли
func ReadLine() (string, error) {
	in := bufio.NewReader(os.Stdin)
	str, err := in.ReadString('\n')
	return strings.Trim(str, "\n"), err
}

package console

import (
	"bufio"
	"os"
)

//Чтение строки из консоли
func ReadLine() (string, error) {
	in := bufio.NewReader(os.Stdin)
	str, err := in.ReadString('\n')
	return str, err
}

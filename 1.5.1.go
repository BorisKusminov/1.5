package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите строку: ")
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Ошибка ввода", err)
		return
	}
	fmt.Println("Количество введенных символов:", len(string(input))-1)
}

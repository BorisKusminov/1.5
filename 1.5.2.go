package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	schetchik := 0
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введите строку: ")
	input, _ := reader.ReadString('\n')
	inputRunesArray := []rune(input)
	for _, v := range inputRunesArray {
		if v == 'а' || v == 'е' || v == 'ё' || v == 'и' || v == 'о' || v == 'у' || v == 'ы' || v == 'э' || v == 'ю' || v == 'я' {
			schetchik++
		}
	}
	fmt.Println(schetchik)
}

package main

import ("bufio" 
"fmt"
	"os"
	"strings"
	"unicode")
	const bukvi string= "яеёиоуыэюя"
func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите строку: ")
	input, err := reader.ReadString('\n')
	if err != nil {
		println("Ошибка ввода", err)
		return
	}
	schetchik :=0
	for _, char := range input {
		lowerChar := unicode.ToLower(char)
		if strings.ContainsRune((bukvi, lolowerChar)){
				schetchik++
						}
		}
			fmt.Println("Количество согласных:", schetchik)
	}


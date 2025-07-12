package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	q := bufio.NewReader(os.Stdin)
	fmt.Println("Введите формулу: ")
	input, _ := q.ReadString('\n')
	result, openBracketCount, closedBacketCount, err := checkCorrectSequenceBracket(input)
	if err != nil {
		fmt.Printf("%s\n", err)
	}
	if result {
		fmt.Printf("Скобки расставлены верно: , %d, открывающиеся, %d закрывающиеся \n", openBracketCount, closedBacketCount)
		fmt.Printf("Скобки расставлены неправильно: , %d, открывающиеся, %d закрывающиеся \n", openBracketCount, closedBacketCount)
	}
}
func checkCorrectSequenceBracket(s string) (bool, int, int, error) {
	var err error = nil
	flag := true
	stack := 0
	runes := []rune(s)
	openBracketCount, closedBacketCount := 0, 0
	for i := range runes {
		if runes[i] == '(' {
			stack++
			openBracketCount++
		} else if runes[i] == ')' {
			closedBacketCount++
			if stack == 0 {
				err = fmt.Errorf("Ставить ')' не закрыв '(' нельзя")
				flag = false
			}
			stack--
		}
	}
	if stack != 0 && flag == true {
		flag = false
	}
	return flag, openBracketCount, closedBacketCount, err
}

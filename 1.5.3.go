package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func capitalizeWords() string {
	q := bufio.NewReader(os.Stdin)
	fmt.Print("Введите текст: ")
	input, _ := q.ReadString('\n') //считаем текст пользователя
	fmt.Println(input)
	words := strings.Fields(input) //разделяем на слова
	result := []string{}           //Создам слайс, в который буду записывать измененные слова
	for _, v := range words {      //идем циклом по каждому слову
		word := []rune(v)                  //каждое слово переводим в руны
		word[0] = unicode.ToUpper(word[0]) //первую руну переводим в верхний регистр
		for i := 1; i < len(word); i++ {
			word[i] = unicode.ToLower(word[i]) // все остальное - в нижний
		}
		result = append(result, string(word)) //добавляю в результирующий слайс переделанное слово
	}
	//Теперь нужно наш слайс result собрать из слайса обратно в предложение и вернуть. Для этого нужно воспользоваться strings
	return strings.Join(result, " ")
}
func main() {
	result := capitalizeWords() //вызываем функцию capitalizeWords. Входных параметров нет, записываем результат в переменную result
	fmt.Println(result)
}

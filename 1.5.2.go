package main

import "fmt"

func main() {
	fmt.Println("Введите текст:")
	var a string
	fmt.Scanln(&a)
	for i := 0; i < len(a); i++ {
		if a[i] == 'а' || a[i] == 'е' || a[i] == 'ё' || a[i] == 'и' || a[i] == 'о' || a[i] == 'у' ||
			a[i] == 'ы' || a[i] == 'э' || a[i] == 'ю' || a[i] == 'я' {
			fmt.Println("Количество согласных:", i)
		}
	}
}

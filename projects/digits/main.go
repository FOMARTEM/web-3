package main

import (
	"fmt"
	"strings"
)

// Дана строка, содержащая только арабские цифры.
// Найти и вывести наибольшую цифру.
func main() {
	str := ""
	fmt.Scan(&str)

	if strings.Contains(str, "9") {
		fmt.Println(9)
	} else if strings.Contains(str, "8") {
		fmt.Println(8)
	} else if strings.Contains(str, "7") {
		fmt.Println(7)
	} else if strings.Contains(str, "6") {
		fmt.Println(6)
	} else if strings.Contains(str, "5") {
		fmt.Println(5)
	} else if strings.Contains(str, "4") {
		fmt.Println(4)
	} else if strings.Contains(str, "3") {
		fmt.Println(3)
	} else if strings.Contains(str, "2") {
		fmt.Println(2)
	} else if strings.Contains(str, "1") {
		fmt.Println(1)
	}
}

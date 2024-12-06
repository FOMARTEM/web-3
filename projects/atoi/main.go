package main

import "fmt"

// На вход подается целое число.
// Необходимо возвести в квадрат каждую цифру числа и вывести получившееся число.
func main() {
	number := ""
	fmt.Scan(&number)
	for i := 0; i < len(number); i++ {
		temp := number[i] - 48
		fmt.Print(temp * temp)
	}
}

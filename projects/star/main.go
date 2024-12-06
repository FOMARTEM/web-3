package main

import "fmt"

func main() {
	str := ""
	fmt.Scan(&str)
	var str2 []rune
	for i := 0; i < len(str)-1; i++ {
		str2 = append(str2, rune(str[i]))
		str2 = append(str2, '*')
	}
	str2 = append(str2, rune(str[len(str)-1]))
	fmt.Println(string(str2))
}

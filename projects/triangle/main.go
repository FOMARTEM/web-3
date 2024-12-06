package main

import (
	"fmt"
	"math"
)

// На вход подаются a и b - катеты прямоугольного треугольника.
// Нужно найти длину гипотенузы
func main() {
	var a, b, c int
	fmt.Scan(&a)
	fmt.Scan(&b)
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}
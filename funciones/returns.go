package main

import "fmt"

func sum(a int, b int) int {
	return a + b
}

// Funciones variadicas

func sum2(values ...int) int {
	var total int
	for _, num := range values {
		total += num
	}

	return total
}

// Retorno de funciones con nombre
func getValues(x int) (double int, triple int, quad int) {
	double = x * 2
	triple = x * 3
	quad = x * 4

	return
}

func main() {
	fmt.Println(sum(1, 3))
	fmt.Println(sum2(1, 3, 4, 5, 6, 7))

	getValues(2)
}

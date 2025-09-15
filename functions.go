package main

import "fmt"

func main() {
	var printValue string = "Hello Go!"
	printMe(printValue)

	numerator, denominator := 3, 2
	var result int = division(numerator, denominator)
	fmt.Println(result)
}

func printMe(printValue string) {
	fmt.Println(printValue)
}

func division(numerator int, denominator int) int {
	var result int = numerator / denominator
	return result
}

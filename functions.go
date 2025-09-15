package main

import (
	"errors"
	"fmt"
)

func main() {
	var printValue string = "Hello Go!"
	printMe(printValue)

	numerator, denominator := 3, 2
	var result, remainder, err = division(numerator, denominator)

	//if-else
	if err != nil {
		fmt.Printf("%s", err.Error())
	} else if remainder == 0 {
		fmt.Printf("The result of the division is %v", result)
	} else {
		fmt.Printf("The result of the integer division is %v with remainder %v", result, remainder)
	}

	//switch-case
	switch {
	case err != nil:
		fmt.Printf("%s", err.Error())
	case remainder == 0:
		fmt.Printf("The result of the division is %v", result)
	default:
		fmt.Printf("The result of the integer division is %v with remainder %v", result, remainder)
	}

	
}

func printMe(printValue string) {
	fmt.Println(printValue)
}

func division(numerator int, denominator int) (int, int, error) {
	var err error

	if denominator == 0 {
		err = errors.New("can't divide by zero")
		return 0, 0, err
	}

	var result int = numerator / denominator
	var remainder int = numerator % denominator

	return result, remainder, err
}

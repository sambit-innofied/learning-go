package main

import (
	"fmt"
	"unicode/utf8"
)

func main(){
	var number int = 3
	fmt.Println(number)

	var floatNum float32 = 3.4
	fmt.Println(floatNum)

	var n1 int = 3
	var n2 int = 2
	fmt.Println(n1/n2)

	var myString string = `Hello World`
	fmt.Println(myString)

	fmt.Println(len("test")) // not the number of characters in the string, it is the number of bytes. 
	fmt.Println(utf8.RuneCountInString("Y")) //length of the string

	var myRune rune = 'a'
	fmt.Println(myRune) //prints the ascii value of the 'a'

}
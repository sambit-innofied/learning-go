package main

import (
	"fmt"
	"unicode/utf8"
)

func mainn(){
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

	var myBoolean bool = true
	fmt.Println(myBoolean)

	// various ways to declare variables
	// var myInt int = 1
	// var myint  = 1
	// myIntt := 1

	myInt, myIntt := 1, 2
	fmt.Println(myInt, myIntt) //myInt = 1, myIntt = 2

	const pi float32 = 3.14
	fmt.Println(pi)


}
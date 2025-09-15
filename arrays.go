package main

import (
	"fmt"
)

func main() {

	// var intArr [3]int32
	// intArr[1] = 123
	// fmt.Println(intArr[0])
	// fmt.Println(intArr[0:2])

	// fmt.Println(&intArr[0])
	// fmt.Println(&intArr[1])

	// var intArr [3]int32 =[3]int32 {1,2,3}
	// intArr := [...]int32{1, 2, 3}
	intArr := [3]int32{1, 2, 3}
	fmt.Println(intArr)

	//slice
	intSlice := []int32{4, 5, 6}
	fmt.Println(intSlice)
	fmt.Printf("The value of the length is %v and capacity is %v", len(intSlice), cap(intSlice))

	intSlice = append(intSlice, 7)
	fmt.Println(intSlice)
	fmt.Printf("The value of the length is %v and capacity is %v", len(intSlice), cap(intSlice))

	var intSlice2 []int32 = []int32{8, 9}
	intSlice = append(intSlice, intSlice2...)
	fmt.Println(intSlice)

	var intSlice3 []int32 = make([]int32, 3, 8)
	fmt.Println(intSlice3)

	//map
	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap)

	var myMap2 map[string]uint8 = map[string]uint8{"adam": 23, "sarah": 45}
	fmt.Println(myMap2["adam"])

	var age, ok = myMap2["adam"]
	if ok {
		fmt.Printf("The age is %v", age)
	} else {
		fmt.Println("Invalid name")
	}
}

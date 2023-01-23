package main

import "fmt"

// golang also supports array
//bit it is a popular practice to use slice instead of array
//slice is similar to array but without a fixed size
//array : var myArray [5]int
//how to accomodate more elements to array after it is fully filled
//you define a new array of double the size of the original
//now, you save the new values in the new array aand delete the old array
//slice : var mySlice []]int

func main() {
	myArray := [5]int{1, 2, 3, 4, 5}

	fmt.Println(myArray)

	/*
		v := "some_string"

		var j string
		j = "some value"
	*/

	mySlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(mySlice)

	mySlice2 := make([]int, 10) // []int{0,0,0,0,0,0,0,0,0,0}

	fmt.Println(mySlice2)

	mySlice3 := make([]int, 0, 10000) // []int{}

	fmt.Println("printing myslice3", mySlice3)
	mySlice4 := make([]int, 0)

	for i := 0; i < 100; i++ {
		mySlice4 = append(mySlice4, 3)
	}

	fmt.Println("printing mySlice4", mySlice4)

	// new function returns the pointer to the object
	// but make function returns the object
	// mySlice5:=new([]int)

	slice5 := mySlice4[:10] // inbox 0,1,2,3,4,5,6,7,8,9

	slice5[4] = 123

	fmt.Println("slice4 after I alter an element of slice5", mySlice4)

	// in case we want to delete the index 4

	slice5 = append(slice5[:4], slice5[5:]...)

	// length  of slice
	l := len(slice5)
	fmt.Println(l)

	for i := 0; i < len(slice5); i++ {
		fmt.Println(slice5[i])
	}

	for index := range slice5 {

		slice5[index] = slice5[index] - 5
		fmt.Println(slice5[index])
	}

	for index, value := range slice5 {
		if index == 0 {
			value = 6
		}
		fmt.Println("index: ", index, " value: ", value)
	}
}

package main

import "fmt"

func main() {

	//integer ; int, unit, int64, int32, uint32, uint64
	//string
	//float ; float64 float32//bool
	//byte
	//rune
	//complex64
	//error : we have a separate data type in golang for defining errors

	//zero value : default value for the data type until any value is given to it
	//integer = 0
	//boolean = false
	//string =""

	//var myData int64

	//fmt.Println(myData)

	//i:=64
	//f:=float64(i)
	//u:=uint(f)

	//strVar:="100"
	//empty var ignores the value
	//intVar, _ := strconv.Atoi(strVar)

	/*
		if err !=nill{
			log.Fatal("kill the program")
		}
	*/

	//stringNumber,err := strconv.Itoa(intVar)

	//constant keyword

	/*
		const data1 = "tasts"
		const data2 = 34
		const data3 = false
	*/

	//for loop with initial condition and final condition
	//and step condition

	/*

		for i := 0; i < 10; i++ {
			//do the operations
		}

		i := 0
		for i < 10 {
			//do the operations
			i = i + 10
		}

		//infinite loop

		for {
			//this is a infinite loop
			if i > 20 {
				break
			}
		}

		if i < 10 {
			fmt.Println("i is less than 10")
		} else {
			fmt.Println("i is greater")
		}

		if v := 10; v < 10 {
			fmt.Println("v is less than 10")
		}
	*/

	code := 10
	switch code {
	case 1:
		//whatever happens at code=1

		fmt.Println("code is equal to 1 inside lower switch case")
		code = 3
		fallthrough
	case 2:
		//whatever happens at code=2
		fmt.Println("code is equal to 2 inside lower switch case")

	case 3:
		//whatever happens at code=2
		fmt.Println("people saying yes won")

	//.....
	case 10:
		fmt.Println("code is equal to 10")
		break
	default:
		fmt.Println("inside the default")
	}
	fmt.Println("the global value of code is", code)

	// type switch

	var v interface{}
	v = 10
	//v= "some_string"
	switch q := v.(type) {
	case bool:
		fmt.Println("this is a boolean")
	case int:
		fmt.Println("this is an integer")
	default:
		fmt.Println("value of type : %T", q)

	}

}

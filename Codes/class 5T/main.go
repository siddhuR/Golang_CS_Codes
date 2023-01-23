package main

import (
	"fmt"
)

/*
funtion : HelloName ()
Input : Name string
Output : "Hello "+ Name
*/

func main() {

	Text := NameString("siddhu")
	fmt.Println(Text)

}

func Print(input string) {
	fmt.Println(input)
}

func NameString(name string) string {
	return "Hello " + name
}

func Nameoutput(name string) (Text string) {

	Text = "Hello " + name
	return
}

func NameVarOutput(names ...string) (Texts []string) {

	return
}

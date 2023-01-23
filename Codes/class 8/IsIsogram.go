package main

/*

Determine if a word or phrase is an isogram.

An isogram (also known as a "non-pattern word") is a word or phrase without a repeating letter, however spaces and hyphens are allowed to appear multiple times.

Examples of isograms:

lumberjacks
background
downstream
six-year-old
The word isograms, however, is not an isogram, because the s repeats.

*/

import "fmt"

var exists = struct{}{}

type set struct {
	m map[string]struct{}
}

func NewSet() *set {

	s := &set{}

	s.m = make(map[string]struct{})

	return s

}

func (s *set) Add(value string) {

	s.m[value] = exists

}

func (s *set) Contains(value string) bool {

	_, c := s.m[value]

	return c

}

func main() {

	s := NewSet()

	var s1 string

	flag := 0

	fmt.Println("Enter a word: ")

	fmt.Scan(&s1)

	for _, character := range s1 {

		if s.Contains(string(character)) {

			fmt.Println("Not an isogram")

			flag = 1

		} else {

			s.Add(string(character))

		}

	}

	if flag == 0 {

		fmt.Println(s1 + " is an isogram")

	}

}

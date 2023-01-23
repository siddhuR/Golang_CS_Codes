package main

import (
	"bufio"

	"fmt"

	"os"

	"strings"

	"github.com/spf13/cast"
)

/*

Alice and Bob went to a pet store. There are �N animals in the store where the ith animal is of type ��Ai​.



Alice decides to buy some of these �N animals. Bob decides that he will buy all the animals left in the store after Alice has made the purchase.



Find out whether it is possible that Alice and Bob end up with exactly same multiset of animals.



Input Format



The first line of input will contain a single integer �T, denoting the number of test cases.

Each test case consists of multiple lines of input.The first line of each test case contains an integer �N — the number of animals in the store.

The next line contains �N space separated integers, denoting the type of each animal.

Output Format



For each test case, output on a new line, YES, if it is possible that Alice and Bob end up with exactly same multiset of animals and NO otherwise.



You may print each character in uppercase or lowercase. For example, the strings YES, yes, Yes, and yES are considered identical.

*/

/*

Testcase 1: 4 4 4 => no

TestCase 2: | 2 3 | 3 2 | => YES

TestCase 3: 1 2 2 3 => NO

TestCase 4: 5 5 1 | 5 1 5

*/

func PetStore() bool {

	fmt.Println("Please enter the input")

	occurence := make(map[int]int)

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n') // "2 3 3 2"

	input = strings.Replace(input, "\n", " ", -1)

	input = strings.TrimSpace(input)

	fmt.Println(" the input is ", input)

	a := strings.Split(input, " ")

	fmt.Println("the array is ", a)

	for _, valStr := range a {

		val := cast.ToInt(valStr)

		fmt.Println(" the input element is ", valStr)

		if count, ok := occurence[val]; ok {

			occurence[val] = count + 1

		} else {

			occurence[val] = 1

		}

	}

	fmt.Println("the map is ", occurence)

	for _, count := range occurence {

		if count%2 != 0 {

			return false

		}

	}

	return true

}

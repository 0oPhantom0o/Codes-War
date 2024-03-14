package main

import (
	"fmt"
	"strconv"
)

/*
	Given an array of integers your solution should find the smallest integer.

For example:

	Given [34, 15, 88, 2] your solution will return 2
	Given [34, -345, -1, 100] your solution will return -345

You can assume, for the purpose of this kata, that the supplied array will not be empty.
*/
//*

func main() {
	//smallestNumber()
	NumberToStirnge()
}

func smallestNumber() {
	number := []int{34, 15, 88, 2, 34, -345, -1, 100}
	emar := []int{1}
	emar[0] = number[0]

	for _, a := range number {
		if a <= emar[0] {
			emar[0] = a
			fmt.Println(emar[0])
		}

	}
	fmt.Println(emar[0])
	//return emar[0]
}

/*
We need a function that can transform a number (integer) into a string.

What ways of achieving this do you know?
Examples (input --> output):

123  --> "123"
999  --> "999"
-100 --> "-100"
*/
func NumberToStirnge() {
	n := 10
	str := strconv.Itoa(n)
	fmt.Println(str)
}

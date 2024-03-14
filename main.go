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
	//NumberToStirnge()
	//BoolToWord()
	//CheckForFactor()
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

/*Complete the method that takes a boolean value and return a "Yes" string for true, or a "No" string for false.*/
func BoolToWord() {
	n := true
	str := strconv.FormatBool(n)
	var a string
	if str == "true" {
		a = "Yes"
	} else {
		a = "No"
	}
	fmt.Println(a)
}

/*
This function should test if the factor is a factor of base.

Return true if it is a factor or false if it is not.
About factors

Factors are numbers you can multiply together to get another number.

2 and 3 are factors of 6 because: 2 * 3 = 6

	You can find a factor by dividing numbers. If the remainder is 0 then the number is a factor.
	You can use the mod operator (%) in most languages to check for a remainder

For example 2 is not a factor of 7 because: 7 % 2 = 1

Note: base is a non-negative number, factor is a positive number.
*/
func CheckForFactor() {
	factor := 3
	base := 9
	var a bool
	if base%factor == 0 {
		a = true
	} else {
		a = false
	}
	fmt.Println(a)
}

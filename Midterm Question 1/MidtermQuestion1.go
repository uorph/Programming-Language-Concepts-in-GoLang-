package main

import "fmt"

// write a function pairLess that takes two slices of int's as parameters
// Pair up the elements and return a slice of bool
// True if the value from the first slice is less than the value from the second slice
// False otherwise
func pairLess(slice1 []int, slice2 []int) []bool {

	result := make([]bool, len(slice1)) //Creating a boolean the size of the arrays
	if len(slice1) != len(slice2) {     // Comparing whether the two arrays are the same size
		fmt.Print("Slices are not the same size, adjust size and try again")
	} else {
		for i := 0; i < len(slice1); i++ { // Iterate for the length of the array
			if slice1[i] < slice2[i] {
				result[i] = true
				//fmt.Print("true", " ")
			} else {
				result[i] = false
				//fmt.Print("false", " ")
			}
		}
	}
	return result
}

func main() {
	slice1 := []int{0, 255, 97, 98}
	slice2 := []int{12, 300, 80, 100}
	answer := pairLess(slice1, slice2)
	for i := 0; i < len(slice1); i++ { //Iterating to print each element of the array
		fmt.Print(answer[i], " ")
	}
}

/*Regina Richardson
CS311*/

package main

import "fmt"

/*
This function returns the calculated pressure using the equation p =nRt/v
Parameter: three float64 variables
Return: a float64 variable which holds the calulated pressure
*/
func CalcPressure(v, n, t float64) (p float64) {
	r := 8.3144598
	return ((r * n * t) / v)
}

// Map for decode
var c = map[byte]byte{
	'e': 'u',
	'h': 'f',
	'l': 'n',
	'o': 'y',
}

/*
	This function decrypts an array using a map

Parameter: array and map, the map decrypts the array
Return: This function returns the decrypted array using the map
*/
func Decode(e []byte, c map[byte]byte) (result3 []byte) {
	for x := 0; x < len(e); x++ {
		var elem byte
		var ok bool

		elem, ok = c[(e[x])]
		if ok {
			e[x] = elem
		}
		// check if key exists in map c
		// if true, change arr[x] to c[arr[x]]
		// else, continue
		// return arr
	}
	return e
}

/*
This function checks for the number of 1s in an array to see if it is an odd number
and also the numbers to determine if it is only 1s and 0s

result1 = true if the list has an odd number of 1s
else false

result2 = true if all the values are 0 or 1
else false

Parameter: An array of integers
Return: two boolean values, depending on the number of odd 1s and whether the array only holds 1s and 0s
*/

func OddParity(arr []int) (result1 bool, result2 bool) {
	var count int
	for x := 0; x < len(arr); x++ {

		if arr[x] == 1 { //Getting the count for the number of 1s
			count++
		}
		// working on getting result2
		if (arr[x] == 1) || (arr[x] == 0) {
			result2 = true
		} else {
			result2 = false
			break
		}
	}

	//working on result 1
	if count%2 == 0 { // count = even
		result1 = false
	} else { // therefore, count is odd
		result1 = true
	}

	return result1, result2

}

func main() {

	fmt.Println(CalcPressure(1.0, 1.0, 298.15))
	fmt.Println(string(Decode([]byte("hello"), c)))
	fmt.Println(string(Decode([]byte("heltastic"), c)))
	fmt.Println(OddParity([]int{1, 1, 1, 0}))
	fmt.Println(OddParity([]int{1, 1, 0, 5}))
}

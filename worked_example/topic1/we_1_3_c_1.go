// "id": "we_1_3_c_1",
// "type": "worked_example",
// "topic": "1. Variables, Data Types, and Operators",
// "difficulty": "hard",
// "context": "Greenhouse Control",
// "title": "Plant Temperature Range Check",
// "description": "Verify greenhouse plant temperature ranges using a new pointer allocation, multi variable threshold declaration, and comparison operators."
package main

import "fmt"

func main() {
	// TODO: Declare a pointer 'tempPtr' of type float64 using the new keyword
	tempPtr := new(float64)
	// TODO: Assign a value of 173.0 to the dereferenced pointer 'tempPtr'
	*tempPtr = 173.0
	// TODO: Declare float64 variable 'minTemp' and assign it a value
	var minTemp float64 = 157.0
	// TODO: Declare float64 variable 'maxTemp' and assign it a value
	var maxTemp float64 = 179.0
	// TODO: Declare a Type Inference variable 'isIdealTemp' (bool) to check if dereferenced pointer 'tempPtr' is greater than 'minTemp' and less than 'maxTemp'
	isIdealTemp := (*tempPtr > minTemp) && (*tempPtr < maxTemp)
	// TODO: Print the value of 'isIdealTemp'
	fmt.Println("Is ideal temp:", isIdealTemp)
}
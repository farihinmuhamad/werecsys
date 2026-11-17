// "id": "we_1_2_d_3",
// "type": "worked_example",
// "topic": "1. Variables, Data Types, and Operators",
// "difficulty": "medium",
// "context": "Chemistry Lab",
// "title": "Acidity Verification",
// "description": "Verify pH level acidity using a new pointer allocation, comparison operators, and string status messages."
package main

import "fmt"

func main() {
	// TODO: Declare a pointer 'phPtr' of type *float64 using the new keyword
	phPtr := new(float64)
	// TODO: Assign the value 6.25 to the dereferenced pointer 'phPtr'
	*phPtr = 6.25
	// TODO: Declare a Type Inference variable 'acidicLimit' (float64) and assign it a value
	acidicLimit := 7.0
	// TODO: Declare a Type Inference variable 'statusAlert' (string) and assign it a value
	statusAlert := "Solution is acidic:"
	// TODO: Declare a Type Inference variable 'isAcidic' (bool) to check if dereferenced pointer 'phPtr' is less than 'acidicLimit'
	isAcidic := *phPtr < acidicLimit
	// TODO: Print the value of statusAlert and isAcidic
	fmt.Println(statusAlert, isAcidic)
}
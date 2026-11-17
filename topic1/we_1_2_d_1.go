// "id": "we_1_2_d_1",
// "type": "worked_example",
// "topic": "1. Variables, Data Types, and Operators",
// "difficulty": "medium",
// "context": "Member Sign-up",
// "title": "Age Verification",
// "description": "Verify client age requirements using a new pointer allocation, comparison operators, and string status messages."
package main

import "fmt"

func main() {
	// TODO: Declare a pointer 'clientAgePtr' of type *int using the new keyword
	clientAgePtr := new(int)
	// TODO: Assign the value 28 to the dereferenced pointer 'clientAgePtr'
	*clientAgePtr = 28
	// TODO: Declare a Type Inference variable 'ageRequirement' (int) and assign it a value
	ageRequirement := 22
	// TODO: Declare a Type Inference variable 'statusReport' (string) and assign it a value
	statusReport := "Age verification status:"
	// TODO: Declare a Type Inference variable 'isQualified' (bool) to check if dereferenced pointer 'clientAgePtr' is greater than 'ageRequirement'
	isQualified := *clientAgePtr > ageRequirement
	fmt.Println(statusReport, isQualified)
}
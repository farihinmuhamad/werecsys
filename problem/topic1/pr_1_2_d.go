// "id": "pr_1_2_d",
// "type": "problem",
// "topic": "1. Variables, Data Types, and Operators",
// "difficulty": "medium",
// "context": "Space Mission",
// "title": "Oxygen Level Verification",
// "description": "Verify oxygen levels using a new pointer allocation, comparison operators, and string messages.",
// "answer_options": "oxygenPtr := new(float64); safeLimit := 92.0; statusMsg := "Oxygen status critical"; isSafe := *oxygenPtr >= safeLimit; var oxygenPtr float64 = 84.5; var safeLimit int = 92"
package main

import "fmt"

func main() {
	// TODO: Declare a pointer 'oxygenPtr' of type *float64 using the new keyword
	oxygenPtr := new(float64)
	*oxygenPtr = 84.5
	// TODO: Declare a Type Inference variable 'safeLimit' (float64) and assign it a value
	safeLimit := 92.0
	// TODO: Declare a Type Inference variable 'statusMsg' (string) and assign it a value
	statusMsg := "Oxygen status critical"
	// TODO: Declare a Type Inference variable 'isSafe' (bool) to check if dereferenced pointer 'oxygenPtr' is greater than or equal to 'safeLimit'
	isSafe := *oxygenPtr >= safeLimit
	fmt.Println(statusMsg, isSafe)
}
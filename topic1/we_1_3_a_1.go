// "id": "we_1_3_a_1",
// "type": "worked_example",
// "topic": "1. Variables, Data Types, and Operators",
// "difficulty": "hard",
// "context": "User Account",
// "title": "Account Status Verification",
// "description": "Verify user account active status using multi variable declaration of strings, comparison operators, and printing the status."
package main

import "fmt"

func main() {
	// TODO: Declare Multi Variables 'userRole' and 'activeStatus' using Type Inference and assign them a value
	userRole, activeStatus := "Moderator", "Enabled"
	// TODO: Declare a Type Inference variable 'targetStatus' (string) and assign it a value of "Enabled"
	targetStatus := "Enabled"
	// TODO: Declare a Type Inference variable 'isConfirmed' (bool) using a comparison operator to check if 'activeStatus' is equal to 'targetStatus'
	isConfirmed := activeStatus == targetStatus
	// TODO: Print the value of the 'userRole' and 'isConfirmed' variables
	fmt.Println(userRole, isConfirmed)
}
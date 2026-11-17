// "id": "we_1_1_b_1",
// "type": "worked_example",
// "topic": "1. Variables, Data Types, and Operators",
// "difficulty": "easy",
// "context": "Buffer Count",
// "title": "Adding Buffer Values",
// "description": "Example showing how to add two integer variables with an intermediate buffer constant."
package main

import "fmt"

func main() {
	// TODO: Declare an integer variable named 'countA' and assign a value to it
	var countA int = 51
	// TODO: Declare an integer variable named 'countB' and assign a value to it
	var countB int = 63
	// TODO: Calculate the sum of 'countA', a buffer constant of 89, and 'countB' into a variable 'totalCount'
	var totalCount int = countA + 89 + countB
	// TODO: Print the value of the 'totalCount' variable
	fmt.Println(totalCount)
}
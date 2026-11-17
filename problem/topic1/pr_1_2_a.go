// "id": "pr_1_2_a",
// "type": "problem",
// "topic": "1. Variables, Data Types, and Operators",
// "difficulty": "medium",
// "context": "Variable swapping",
// "title": "Swap Two Integer Variables",
// "description": "Swap the values of two integer variables using a temporary variable."
// "answer_options": "var swapNumA int = 125; swapTemp := swapNumA; swapNumA = swapNumB; swapNumB = swapTemp; var swapTemp int; var swapNumB int = 80"
package main

import "fmt"

func main() {
	// TODO: Declare an integer variable named 'swapNumA' and assign a value to it
	var swapNumA int = 125
	var swapNumB int = 80
	fmt.Println("Before swap:", swapNumA, swapNumB)
	// TODO: Declare a Type Inference variable 'swapTemp' and assign 'swapNumA' to it
	swapTemp := swapNumA
	// TODO: Assign the value of 'swapNumB' to 'swapNumA'
	swapNumA = swapNumB
	// TODO: Assign the value of 'swapTemp' to 'swapNumB'
	swapNumB = swapTemp
	fmt.Println("After swap:", swapNumA, swapNumB)
}
// "id": "pr_2_1_d",
// "type": "problem",
// "topic": "2. Branching (Conditional Statements)",
// "difficulty": "easy",
// "context": "Number Classification",
// "title": "Check if a Number is Even or Odd",
// "description": "In this problem, you will learn how to check if a number is even or odd using the switch statement and modulus operator. Declare an integer variable, assign it a value, and then use a switch statement to determine if the number is even or odd based on the result of the modulus operation.",
// "answer_options": "var inputNumber int = 7; switch inputNumber % 2 {; case 0:; fmt.Println("Odd"); switch inputNumber {; case inputNumber:"
package main

import "fmt"

func main() {
	// TODO: Declare an integer variable 'inputNumber' and assign it a value
	var inputNumber int = 7
	// TODO: Use switch statement to check inputNumber with modulus operation
	switch inputNumber % 2 {
	// TODO: Case statement to check for even number (remainder 0)
	case 0:
		fmt.Println("Even")
	default:
		// TODO: Print "Odd" for default case (remainder 1)
		fmt.Println("Odd")
	}
}
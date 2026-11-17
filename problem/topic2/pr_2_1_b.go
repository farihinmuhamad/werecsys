// "id": "pr_2_1_b",
// "type": "problem",
// "topic": "2. Branching (Conditional Statements)",
// "difficulty": "easy",
// "context": "Check Password",
// "title": "Password Validation",
// "description": "In this problem, you will create a simple password validation program. The program will check if the user's input matches a predefined password and will print an appropriate message based on the result. If the input string matches a predefined password, print 'Access Granted'. Otherwise, print 'Access Denied'.",
// "answer_options": "password := "secret123"; var input string = "secret"; if input == password {; } else {; { else }; cons input string = "secret""
package main

import "fmt"

func main() {
	// TODO: Declare a Type Inference variable 'password' and assign it a value
	password := "secret123"
	// TODO: Declare a variable 'input' of type string and assign it a value
	var input string = "secret"
	// TODO: Create an if statement to check if 'input' is equal to 'password'
	if input == password {
		fmt.Println("Access Granted")
	// TODO: Add an else block to handle if the input does not match the password
	} else {
		fmt.Println("Access Denied")
	}
}
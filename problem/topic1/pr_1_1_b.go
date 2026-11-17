// "id": "pr_1_1_b",
// "type": "problem",
// "topic": "1. Variables, Data Types, and Operators",
// "difficulty": "easy",
// "context": "String Manipulation",
// "title": "Full Name Concatenation",
// "description": "Create a program that joins first name and last name into a full name."
// "answer_options": "var firstName string = "Billy"; var lastName string = "Joe"; var fullName string = firstName + " " + lastName; fmt.Println(fullName); fullName := firstName; var firstName int = 10"
package main

import "fmt"

func main() {
	// TODO: Declare a string variable 'firstName' and assign it with with a name
	var firstName string = "Billy"
	// TODO: Declare a string variable 'lastName' and assign it with with a name
	var lastName string = "Joe"
	// TODO: Concatenate 'firstName', a space, and 'lastName' into a variable 'fullName'
	var fullName string = firstName + " " + lastName
	// TODO: Print the value of the 'fullName' variable
	fmt.Println(fullName)
}
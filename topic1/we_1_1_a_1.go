// "id": "we_1_1_a_1",
// "type": "worked_example",
// "topic": "1. Variables, Data Types, and Operators",
// "difficulty": "easy",
// "context": "String Concatenation",
// "title": "Concatenating Names",
// "description": "Example showing how to concatenate two string variables."
package main

import "fmt"

func main() {
	// TODO: Declare a string variable named 'givenName' and assign a value to it
	var givenName string = "Grace"
	// TODO: Declare a string variable named 'surName' and assign a value to it
	var surName string = "Hopper"
	// TODO: Calculate the concatenation of 'givenName' and 'surName' and store it in a new variable 'completeName'
	var completeName string = givenName + surName
	// TODO: Print the value of the 'completeName' variable
	fmt.Println(completeName)
}
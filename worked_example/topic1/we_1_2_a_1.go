// "id": "we_1_2_a_1",
// "type": "worked_example",
// "topic": "1. Variables, Data Types, and Operators",
// "difficulty": "medium",
// "context": "Job Titles",
// "title": "Swapping Job Titles",
// "description": "Swap two string variables representing job titles using a temporary variable."
package main

import "fmt"

func main() {
	// TODO: Declare a string variable named 'firstTitle' and assign a value to it
	var firstTitle string = "Director"
	// TODO: Declare a string variable named 'secondTitle' and assign a value to it
	var secondTitle string = "Designer"
	fmt.Println("Before swap:", firstTitle, secondTitle)
	// TODO: Declare a Type Inference variable 'tempTitle' and assign 'firstTitle' to it
	tempTitle := firstTitle
	// TODO: Assign the value of 'secondTitle' to 'firstTitle'
	firstTitle = secondTitle
	// TODO: Assign the value of 'tempTitle' to 'secondTitle'
	secondTitle = tempTitle
	// TODO: Print the values of 'firstTitle' and 'secondTitle' variables
	fmt.Println("After swap:", firstTitle, secondTitle)
}
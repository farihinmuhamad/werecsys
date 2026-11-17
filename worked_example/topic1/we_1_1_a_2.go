// "id": "we_1_1_a_2",
// "type": "worked_example",
// "topic": "1. Variables, Data Types, and Operators",
// "difficulty": "easy",
// "context": "Score Addition",
// "title": "Adding Game Scores",
// "description": "Example showing how to add two integer variables using multi-variable manifest declaration."
package main

import "fmt"

func main() {
	// TODO: Declare integer variables 'scoreA' and 'scoreB' using multi-variable manifest declaration and assign them values
	var scoreA, scoreB int = 48, 62
	// TODO: Calculate the sum of 'scoreA' and 'scoreB' and store it in a new variable 'totalScore'
	var totalScore int = scoreA + scoreB
	// TODO: Print the value of the 'totalScore' variable
	fmt.Println(totalScore)
}
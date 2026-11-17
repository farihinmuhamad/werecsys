// "id": "we_1_2_a_2",
// "type": "worked_example",
// "topic": "1. Variables, Data Types, and Operators",
// "difficulty": "medium",
// "context": "Score Swapping",
// "title": "Swapping Game Scores",
// "description": "Swap two game score integer variables without using a temporary variable."
package main

import "fmt"

func main() {
	// TODO: Declare type inference variable 'playerAScore' and assign it a value
	playerAScore := 340
	// TODO: Declare type inference variable 'playerBScore' and assign it a value
	playerBScore := 210
	fmt.Println("Initial scores:", playerAScore, playerBScore)
	// TODO: Swap 'playerAScore' and 'playerBScore' in a single line using tuple assignment
	playerAScore, playerBScore = playerBScore, playerAScore
	// TODO: Print the values of 'playerAScore' and 'playerBScore' variables
	fmt.Println("Swapped scores:", playerAScore, playerBScore)
}
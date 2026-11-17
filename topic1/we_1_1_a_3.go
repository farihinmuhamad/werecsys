// "id": "we_1_1_a_3",
// "type": "worked_example",
// "topic": "1. Variables, Data Types, and Operators",
// "difficulty": "easy",
// "context": "Inventory Count",
// "title": "Summing Inventory Items",
// "description": "Example showing how to add two integer variables representing inventory counts."
package main

import "fmt"

func main() {
	// TODO: Declare an integer variable named 'boxCount' and assign a value to it
	var boxCount int = 79
	// TODO: Declare an integer variable named 'crateCount' and assign a value to it
	var crateCount int = 91
	// TODO: Calculate the sum of 'boxCount' and 'crateCount' and store it in a new variable 'totalStorage'
	var totalStorage int = boxCount + crateCount
	// TODO: Print the value of the 'totalStorage' variable
	fmt.Println("Total Storage:", totalStorage)
}
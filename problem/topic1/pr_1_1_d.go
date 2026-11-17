// "id": "pr_1_1_d",
// "type": "problem",
// "topic": "1. Variables, Data Types, and Operators",
// "difficulty": "easy",
// "context": "Status Check",
// "title": "Switch Status",
// "description": "Store a boolean status and print its inverse (NOT operator)."
// "answer_options": "var isOn bool = true; fmt.Println(isOn); var isOff bool = !isOn; fmt.Println(isOff); isOff := isOn; var isOn bool = false"
package main

import "fmt"

func main() {
	// TODO: Declare a boolean variable 'isOn' and set it to true
	var isOn bool = true
	// TODO: Print the value of the 'isOn' variable
	fmt.Println(isOn)
	// TODO: Declare a boolean variable 'isOff' which is the negation of 'isOn'
	var isOff bool = !isOn
	// TODO: Print the value of the 'isOff' variable
	fmt.Println(isOff)
}
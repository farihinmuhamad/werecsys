// "id": "pr_1_1_c",
// "type": "problem",
// "topic": "1. Variables, Data Types, and Operators",
// "difficulty": "easy",
// "context": "Area Calculation",
// "title": "Calculate Area of a Rectangle",
// "description": "Multiply two float values representing length and width to calculate area."
// "answer_options": "var length float64 = 10.5; var width float64 = 2.4; var area float64 = length * width; fmt.Println(area); area := length + width; var width float64 = length"
package main

import "fmt"

func main() {
	// TODO: Declare a float64 Variable 'length' and assign it a value
	var length float64 = 10.5
	// TODO: Declare a float64 Variable 'width' and assign it a value
	var width float64 = 2.4
	// TODO: Multiply length and width, then store it in a new float64 variable 'area'
	var area float64 = length * width
	// TODO: Print the value of the 'area' variable
	fmt.Println(area)
}
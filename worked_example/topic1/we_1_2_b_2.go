// "id": "we_1_2_b_2",
// "type": "worked_example",
// "topic": "1. Variables, Data Types, and Operators",
// "difficulty": "medium",
// "context": "Circular Measure",
// "title": "Circle Perimeter with Diameter",
// "description": "Calculate the perimeter of a circle using diameter and Pi constant."
package main

import "fmt"

func main() {
	// TODO: Declare a constant 'PiVal' of type float64 and assign it a value
	const PiVal float64 = 3.14159265
	// TODO: Declare a float64 variable 'diameter' and assign it a value
	var diameter float64 = 11.2
	// TODO: Calculate the circle perimeter and store it in a float64 variable 'perimeter'
	var perimeter float64 = PiVal * diameter
	// TODO: Print the value of the 'perimeter' variable
	fmt.Println(perimeter)
}
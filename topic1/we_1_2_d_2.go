// "id": "we_1_2_d_2",
// "type": "worked_example",
// "topic": "1. Variables, Data Types, and Operators",
// "difficulty": "medium",
// "context": "Pressure Check",
// "title": "Boiler Pressure Monitor",
// "description": "Monitor boiler pressure using float64 pointers, dereferencing comparison, and manifest variable declarations."
package main

import "fmt"

func main() {
	// TODO: Declare a float64 variable 'currentPressure' and assign a value to it
	var currentPressure float64 = 14.2
	// TODO: Declare a float64 pointer variable 'pressurePtr' and assign the address of 'currentPressure' to it
	var pressurePtr *float64 = &currentPressure
	// TODO: Declare a manifest float64 variable 'maxThreshold' and assign a value to it
	var maxThreshold float64 = 16.5
	// TODO: Declare a manifest string variable 'boilerAlert' and assign a value to it
	var boilerAlert string = "Pressure within bounds:"
	// TODO: Declare a manifest boolean variable 'isNormal' checking if dereferenced pointer 'pressurePtr' is less than or equal to 'maxThreshold'
	var isNormal bool = *pressurePtr <= maxThreshold
	// TODO: Print the value of boilerAlert and isNormal
	fmt.Println(boilerAlert, isNormal)
}
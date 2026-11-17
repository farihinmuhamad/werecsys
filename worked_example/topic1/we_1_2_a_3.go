// "id": "we_1_2_a_3",
// "type": "worked_example",
// "topic": "1. Variables, Data Types, and Operators",
// "difficulty": "medium",
// "context": "Velocity Adjustment",
// "title": "Swapping Vehicle Speeds",
// "description": "Swap two vehicle speed integer variables using a temporary variable."
package main

import "fmt"

func main() {
	// TODO: Declare an integer variable named 'speedA' and assign a value to it
	var speedA int = 185
	// TODO: Declare an integer variable named 'speedB' and assign a value to it
	var speedB int = 65
	// TODO: Print the values of 'speedA' and 'speedB' variables before swap
	fmt.Println(speedA, speedB)
	// TODO: Declare a Type Inference variable 'speedTemp' and assign 'speedA' to it
	speedTemp := speedA
	// TODO: Assign the value of 'speedB' to 'speedA'
	speedA = speedB
	// TODO: Assign the value of 'speedTemp' to 'speedB'
	speedB = speedTemp
	// TODO: Print the values of 'speedA' and 'speedB' variables after swap
	fmt.Println(speedA, speedB)
}
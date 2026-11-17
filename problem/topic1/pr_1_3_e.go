// "id": "pr_1_3_e",
// "type": "problem",
// "topic": "1. Variables, Data Types, and Operators",
// "difficulty": "hard",
// "context": "Petrochemical Processing",
// "title": "Process Pressure Safety Check",
// "description": "Check if the process pressure readings are within safe limits. Declare pressure readings, define a critical limit using a constant, and evaluate safety status using logical operators.",
// "answer_options": "const CriticalLimit int = 620; var systemMode string = "Auto"; sensorVal := 495; isSafeState := (sensorVal < CriticalLimit) && (systemMode == "Auto"); CriticalLimit int = 620; systemMode string = "Auto""
package main

import "fmt"

func main() {
	// TODO: Declare an integer constant 'CriticalLimit' and assign it a value
	const CriticalLimit int = 620
	// TODO: Declare a string variable 'systemMode' using and assign it a value
	var systemMode string = "Auto"
	// TODO: Declare an integer variable 'sensorVal' using Type Inference
	sensorVal := 495
	// TODO: Declare a Type Inference variable 'isSafeState' (bool) using comparison and logical operators to check if sensor values are below 'CriticalLimit' AND 'systemMode' equals "Auto"
	isSafeState := (sensorVal < CriticalLimit) && (systemMode == "Auto")
	fmt.Println(isSafeState)
}
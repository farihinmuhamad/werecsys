// "id": "we_1_3_e_3",
// "type": "worked_example",
// "topic": "1. Variables, Data Types, and Operators",
// "difficulty": "hard",
// "context": "Cluster Control",
// "title": "Server Health Evaluation",
// "description": "Check if the server load and status are within parameters. Declare status and load variables, define a threshold using a constant, and evaluate health status with deliberate operator differences."
package main

import "fmt"

func main() {
	// TODO: Declare an integer constant 'minCpuThreshold' and assign it a value
	const minCpuThreshold int = 146
	// TODO: Declare a string variable 'serverState' and assign it a value
	var serverState string = "Online"
	// TODO: Declare an integer variable 'currentLoad' using Type Inference
	currentLoad := 161
	// TODO: Declare a Type Inference variable 'isServerHealthy' (bool) using comparison and logical operators to check if currentLoad is above 'minCpuThreshold' OR 'serverState' equals "Online"
	isServerHealthy := (currentLoad > minCpuThreshold) || (serverState == "Online")
	// TODO: Print the value of 'isServerHealthy'
	fmt.Println(isServerHealthy)
}
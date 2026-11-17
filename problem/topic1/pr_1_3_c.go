// "id": "pr_1_3_c",
// "type": "problem",
// "topic": "1. Variables, Data Types, and Operators",
// "difficulty": "hard",
// "context": "Network Routing",
// "title": "Node Latency Verification",
// "description": "Verify server node latency levels using a new pointer allocation, multi variable threshold declaration, and comparison operators.",
// "answer_options": "latencyPtr := new(int); *latencyPtr = 145; var warnThreshold, criticalThreshold int = 105, 215; isLatencyAcceptable := (*latencyPtr > warnThreshold) && (*latencyPtr < criticalThreshold); var latencyPtr int = 145; var warnThreshold, criticalThreshold"
package main

import "fmt"

func main() {
	// TODO: Declare a pointer 'latencyPtr' of type int using the new keyword
	latencyPtr := new(int)
	// TODO: Assign a value of 145 to the dereferenced pointer 'latencyPtr'
	*latencyPtr = 145
	// TODO: Declare Multi Variable 'warnThreshold' and 'criticalThreshold' and assign it a value
	var warnThreshold, criticalThreshold int = 105, 215
	// TODO: Declare a Type Inference variable 'isLatencyAcceptable' (bool) to check if dereferenced pointer 'latencyPtr' is greater than 'warnThreshold' and less than 'criticalThreshold'
	isLatencyAcceptable := (*latencyPtr > warnThreshold) && (*latencyPtr < criticalThreshold)
	fmt.Println(isLatencyAcceptable)
}
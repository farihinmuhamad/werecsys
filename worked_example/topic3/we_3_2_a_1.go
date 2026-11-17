// "id": "we_3_2_a_1",
// "type": "worked_example",
// "topic": "3. Looping",
// "difficulty": "medium",
// "context": "Pipeline Monitoring",
// "title": "Flow Count Accumulator",
// "description": "Accumulate flow event counts over multiple monitoring cycles using a regular for loop, integer types, and short variable declarations."
package main

import "fmt"

func main() {
	// TODO: Declare an integer variable 'flowIncrement' and assign it a value
	var flowIncrement int = 3
	// TODO: Declare an integer variable 'monitorCycles' and assign it a value
	var monitorCycles int = 16
	// TODO: Declare an integer variable 'totalFlowCount' and assign it a value
	var totalFlowCount int = 0
	// TODO: Create a regular for loop to iterate 'counter' from 1 to 'monitorCycles', increment 'counter' by 1 at each iteration
	for counter := 1; counter <= monitorCycles; counter++ {
		// TODO: Change the value of 'totalFlowCount' by adding 'flowIncrement' to it
		totalFlowCount += flowIncrement
	}
	// TODO: Print the value of 'totalFlowCount' to the console
	fmt.Println(totalFlowCount)
}
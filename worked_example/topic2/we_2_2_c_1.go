// "id": "we_2_2_c_1",
// "type": "worked_example",
// "topic": "2. Branching (Conditional Statements)",
// "difficulty": "medium",
// "context": "System Resource Monitor",
// "title": "Processor Load Status",
// "description": "Assess server processor and memory load levels. Declare two float64 variables using multi-variable declaration and determine the load condition."
package main

import "fmt"

func main() {
	// TODO: Declare a constant 'ProcessorLimit' of type float64 and assign it a value
	const ProcessorLimit float64 = 95.0
	// TODO: Declare two float64 variables 'processorLoad' and 'memoryLoad' using Multi Variable Declaration and assign them values
	var processorLoad, memoryLoad float64 = 85.5, 45.7
	// TODO: Create an if statement to check if 'processorLoad' is greater than 'ProcessorLimit' AND 'memoryLoad' is greater than 90.0
	if processorLoad > ProcessorLimit && memoryLoad > 90.0 {
		// TODO: Print the "System Critical" status message
		fmt.Println("System Critical")
		// TODO: Create an else-if statement to check if 'processorLoad' is greater than 75.0 AND 'memoryLoad' is greater than 40.0
	} else if processorLoad > 75.0 && memoryLoad > 40.0 {
		// TODO: Print the "System Warning" status message
		fmt.Println("System Warning")
		// TODO: Else statement to handle normal priority status
	} else {
		// TODO: Print the "System Normal" status message if none of the above conditions are met
		fmt.Println("System Normal")
	}
}

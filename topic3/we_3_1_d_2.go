// "id": "we_3_1_d_2",
// "type": "worked_example",
// "topic": "3. Looping",
// "difficulty": "easy",
// "context": "Submersible Flow System",
// "title": "Pump Flow Integration",
// "description": "Integrate total pumped water volume across multiple scheduled run cycles using a regular 3-clause loop."
package main

import "fmt"

func main() {
	// TODO: Declare integer variables 'totalPumpRuns' and assign initial values
	var totalPumpRuns int = 5
	// TODO: Declare integer variables 'flowRatePerRun' and assign initial values
	var flowRatePerRun int = 15
	// TODO: Declare an integer variable 'totalPumped' and assign initial values
	var totalPumped int = 0
	// TODO: Create a regular for loop to iterate over the pump run cycles
	for cycle := 1; cycle <= totalPumpRuns; cycle++ {
		// TODO: Increment 'totalPumped' by 'flowRatePerRun' for each cycle
		totalPumped += flowRatePerRun
		// TODO: Print the current cycle number
		fmt.Println("Cycle completed:", cycle)
	}
	// TODO: Print the total pumped volume after the loop
	fmt.Println(totalPumped)
}

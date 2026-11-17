// "id": "we_3_1_c_1",
// "type": "worked_example",
// "topic": "3. Looping",
// "difficulty": "easy",
// "context": "Cluster Allocation",
// "title": "Cluster Capacity Calculator",
// "description": "Calculate consecutive cluster capacities along expansion steps using a loop, constants, and short variable declarations."
package main

import "fmt"

func main() {
	// TODO: Declare an integer constant 'IncrementStep' and assign it a value
	const IncrementStep int = 3
	// TODO: Declare a string variable 'clusterNodeName' and assign it a value using short declaration
	clusterNodeName := "Node-X"
	// TODO: Declare an integer variable 'capacityUnits' and assign it a value
	var capacityUnits int = 6
	// TODO: Create a regular for loop to iterate 'i' from 1 to 3 while incrementing 'i' by 1 at each iteration
	for i := 1; i <= 3; i++ {
		// TODO: Add the value of 'IncrementStep' to 'capacityUnits'
		capacityUnits += IncrementStep
	}
	// TODO: Print the values of 'clusterNodeName' and 'capacityUnits'
	fmt.Println(clusterNodeName, capacityUnits)
}
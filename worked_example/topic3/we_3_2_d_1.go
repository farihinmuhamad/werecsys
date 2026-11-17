// "id": "we_3_2_d_1",
// "type": "worked_example",
// "topic": "3. Looping",
// "difficulty": "medium",
// "context": "Port Scanner",
// "title": "Port Scanner Status",
// "description": "Scan network ports, print status message for specific indices using continue only without break, using Type Inference Variable Declarations. Declare variables for status message, target port, and maximum ports. Iterate through port indices, print status for the target port, and skip others."
package main

import "fmt"

func main() {
	// TODO: Declare a Type Inference variable 'portStatus' (string) and assign it a value
	portStatus := "Open"
	// TODO: Declare a Type Inference variable 'targetPort' (int) and assign it a value
	targetPort := 132
	// TODO: Declare a Type Inference variable 'maxPorts' (int) and assign it a value
	maxPorts := 138
	// TODO: Create a regular for loop to iterate 'portIndex' from 1 to 'maxPorts' with increment of 1 at each iteration
	for portIndex := 1; portIndex <= maxPorts; portIndex++ {
		// TODO: Create an if statement to check if 'portIndex' is less than 'targetPort'
		if portIndex < targetPort {
			// TODO: Use 'continue' statement to skip the current iteration to the next iteration of the loop
			continue
		}
		// TODO: Print 'portStatus' and the current index 'portIndex' when 'portIndex' is equal to 'targetPort'
		fmt.Println(portStatus, portIndex)
	}
}
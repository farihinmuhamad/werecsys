// "id": "we_3_3_b_2",
// "type": "worked_example",
// "topic": "3. Looping",
// "difficulty": "hard",
// "context": "Network Ping",
// "title": "Request Tracker",
// "description": "Monitor network requests by counting pings up to a limit, using a standard 3-clause for loop, constants, and explicit variable declarations. Note: conceptual similarity applied — uses a standard 3-clause loop instead of a while-like loop, and explicit variable declarations.",

package main

import "fmt"

func main() {
	// TODO: Declare a constant integer variable 'routeId' and assign it a value
	const routeId int = 847
	// TODO: Declare a constant boolean variable 'checkActive' and assign it a value
	const checkActive bool = true
	// TODO: Declare an integer variable 'pingCount' and assign it a value
	var pingCount int = 0
	// TODO: Declare an integer variable 'maxPings' and assign it a value
	var maxPings int = 101
	// TODO: Create a standard 3-clause for loop to count pings from 1 to maxPings
	for i := 1; i <= maxPings; i++ {
		// TODO: Increment 'pingCount' inside the loop
		pingCount++
	}
	// TODO: Print the values of 'routeId', 'checkActive', and 'pingCount'
	fmt.Println(routeId, checkActive, pingCount)
}
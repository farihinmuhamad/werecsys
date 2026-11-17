// "id": "pr_3_2_c",
// "type": "problem",
// "topic": "3. Looping",
// "difficulty": "medium",
// "context": "Telemetry Stream",
// "title": "Stream Packet Parser",
// "description": "Process incoming telemetry packets in an infinite loop. Print the packet marker and the index, increment the counter, and break when a target index is reached using Type Inference Variable Declarations. Declare and initialize variables for packet marker, target packet index, and current packet index. Use a for loop without arguments to create an infinite loop, and implement a break condition to exit the loop when the target packet is reached.",
// "answer_options": "for {; if currentPacket == targetPacket {; break; currentPacket++; for packetMarker == "Beacon" {; while currentPacket = targetPacket {"
package main

import "fmt"

func main() {
	packetMarker := "Beacon"
	targetPacket := 61
	currentPacket := 1
	// TODO: Create a for loop without any arguments (infinite loop)
	for {
		fmt.Println(packetMarker, currentPacket)
		// TODO: Create an if statement to check if 'currentPacket' is equal to 'targetPacket'
		if currentPacket == targetPacket {
			// TODO: Use 'break' statement to exit the loop when target packet is reached
			break
		}
		// TODO: Increment 'currentPacket' by 1 manually inside the loop
		currentPacket++
	}
}
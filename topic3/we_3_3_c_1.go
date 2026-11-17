// "id": "we_3_3_c_1",
// "type": "worked_example",
// "topic": "3. Looping",
// "difficulty": "hard",
// "context": "Signal Processing",
// "title": "Target Packet Scanner",
// "description": "Search a series of packets for a critical signal. Print a search message, loop through iterations, skip normal packets with continue, and stop when a critical packet is found using break. Declare variables for scan action, packet limit, and critical packet; use a for loop to iterate through packets, print scan action and packet number, use continue to skip normal packets, and break when the critical packet is reached.",

package main

import "fmt"

func main() {
	// TODO: Declare a string variable 'scanAction' and assign it a value
	var scanAction string = "Isolate"
	// TODO: Declare two integer variables 'packetLimit' and 'criticalPacket' using multi variable declaration and assign them values
	var packetLimit, criticalPacket int = 159, 158
	// TODO: Create a regular for loop to iterate 'packetIndex' from 1 to 'packetLimit' while incrementing 'packetIndex' by 1 at each iteration
	for packetIndex := 1; packetIndex <= packetLimit; packetIndex++ {
		// TODO: Create an if statement to check if 'packetIndex' is less than 'criticalPacket'
		if packetIndex < criticalPacket {
			continue
		}
		fmt.Println(scanAction, packetIndex)
		// TODO: Create an if statement to check if 'packetIndex' is equal to 'criticalPacket'
		if packetIndex == criticalPacket {
			// TODO: Execute 'break' statement to exit the loop when 'packetIndex' is equal to 'criticalPacket'
			break
		}
	}
}

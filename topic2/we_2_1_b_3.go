// "id": "we_2_1_b_3",
// "type": "worked_example",
// "topic": "2. Branching (Conditional Statements)",
// "difficulty": "easy",
// "context": "Data Security",
// "title": "Hash Code Match Check",
// "description": "Verify integrity of downloaded files. Declare target and user hashes, then check if they do not match with deliberate operator difference."
package main

import "fmt"

func main() {
	// TODO: Declare a Type Inference variable 'targetHash' and assign it a value
	targetHash := "hash8732"
	// TODO: Declare a variable 'userHash' of type string and assign it a value
	var userHash string = "hash8730"
	// TODO: Create an if statement to check if 'userHash' is not equal to 'targetHash'
	if userHash != targetHash {
		// TODO: Inside the if block, print the value of "Key Mismatch"
		fmt.Println("Key Mismatch")
		// TODO: Add an else block
	} else {
		// TODO: Inside the else block, print the value of "Key Match"
		fmt.Println("Key Match")
	}
}

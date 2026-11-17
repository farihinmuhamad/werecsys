// "id": "we_2_3_c_2",
// "type": "worked_example",
// "topic": "2. Branching (Conditional Statements)",
// "difficulty": "hard",
// "context": "Hospital Triage System",
// "title": "Evaluate Hospital Triage Priority",
// "description": "Determine ICU triage priority status based on patient heart rate, asthma history, and urgent case flag."
package main

import "fmt"

func main() {
	// TODO: create three variables: isUrgentCase (bool), heartRate (int), and hasAsthma (bool) and assign them appropriate values
	var isUrgentCase bool = false
	var heartRate int = 135
	var hasAsthma bool = false
	// TODO: Create an if statement to check if isUrgentCase is true and heartRate is greater than 120
	if isUrgentCase && heartRate > 120 {
		// TODO: Create another if statement to check if hasAsthma is true
		if hasAsthma {
			fmt.Println("Urgent ICU Priority - Asthma Care")
		} else {
			fmt.Println("Urgent ICU Priority Assigned")
		}
		// TODO: Create an else-if statement to check if isUrgentCase is false and heartRate is greater than 150
	} else if !isUrgentCase && heartRate > 150 {
		// TODO: Create another if statement to check if hasAsthma is true
		if hasAsthma {
			fmt.Println("Standard ICU Priority - Asthma Care")
		} else {
			fmt.Println("Standard ICU Priority Assigned")
		}
		// TODO: Create an else statement to print Routine Ward Assignment if none of the conditions are met
	} else {
		fmt.Println("Routine Ward Assignment")
	}
}

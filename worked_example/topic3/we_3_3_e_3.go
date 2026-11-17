// "id": "we_3_3_e_3",
// "type": "worked_example",
// "topic": "3. Looping",
// "difficulty": "hard",
// "context": "Memory Map",
// "title": "Led Matrix Mapper",
// "description": "Iterate through rows and columns of an LED controller grid in reverse. Print the LED icon and cell coordinates using nested decrementing for loops and short variable declarations. Declare variables for LED icon, LED tag, total rows, and total columns. Note: deliberate differences applied — uses nested decrementing loops and short variable declarations.",

package main

import "fmt"

func main() {
	// TODO: Declare string variables 'ledIcon' and 'ledTag' using short variable declarations and assign them values
	ledIcon := "~"
	ledTag := "LED"
	// TODO: Declare a type inference variable 'totalRows' (int) and 'totalCols' (int) and assign them values using short variable declarations
	totalRows, totalCols := 192, 193
	// TODO: Create outer decrementing for loop to iterate 'r' from 'totalRows' down to 1
	for r := totalRows; r >= 1; r-- {
		// TODO: Create inner decrementing for loop to iterate 'c' from 'totalCols' down to 1
		for c := totalCols; c >= 1; c-- {
			// TODO: Print the values of 'ledIcon', 'ledTag', 'r', and 'c'
			fmt.Println(ledIcon, ledTag, r, c)
		}
	}
}
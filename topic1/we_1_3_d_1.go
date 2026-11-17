// "id": "we_1_3_d_1",
// "type": "worked_example",
// "topic": "1. Variables, Data Types, and Operators",
// "difficulty": "hard",
// "context": "Shipping Cart",
// "title": "Shipping Checkout Validation",
// "description": "Determine checkout shipping qualification using constants, multi-variable boolean declaration, and logical evaluation."
package main

import "fmt"

func main() {
	// TODO: Declare type inference int variable 'weight' and assign it a value
	weight := 241
	// TODO: Declare 'minimumWeight' as a constant int and assignt it a value
	const minimumWeight int = 229
	// TODO: Declare a boolean variable 'freeShipping', initialize it with true
	var freeShipping bool = true
	// TODO: Declare a boolean variable 'isPromoZone', initialize it with false
	var isPromoZone bool = false
	// TODO: Declare Multi Variable 'checkWeight', 'checkPromo', and 'shipEligible' as boolean
	var checkWeight, checkPromo, shipEligible bool
	// TODO: Determine 'checkWeight': True if weight is greater than or equal to 'minimumWeight'
	checkWeight = (weight >= minimumWeight)
	// TODO: Determine 'checkPromo': True if 'freeShipping' is true OR 'isPromoZone' is true
	checkPromo = (freeShipping || isPromoZone)
	// TODO: Determine 'shipEligible': Shipping is qualified if 'checkWeight' is true AND 'checkPromo' is true
	shipEligible = checkWeight && checkPromo
	// TODO: Print the value of 'shipEligible'
	fmt.Println(shipEligible)
}
// "id": "we_1_1_d_3",
// "type": "worked_example",
// "topic": "1. Variables, Data Types, and Operators",
// "difficulty": "easy",
// "context": "Subscription Status",
// "title": "Negating Subscription Status",
// "description": "Example showing how to negating a subscription status using manifest typing."
package main

import "fmt"

func main() {
	// TODO: Declare a boolean variable 'isSubscribed' and set it to true
	var isSubscribed bool = true
	// TODO: Print the value of the 'isSubscribed' variable
	fmt.Println(isSubscribed)
	// TODO: Declare a boolean variable 'isUnsubscribed' which is the negation of 'isSubscribed'
	isUnsubscribed := !isSubscribed
	// TODO: Print the value of the 'isUnsubscribed' variable
	fmt.Println(isUnsubscribed)
}
package main

import (
	"fmt"
)

func main() {
	// Variable to hold the user input
	var number int

	// Prompt the user to enter a number
	fmt.Print("Enter a number: ")
	fmt.Scan(&number)

	// If-Else statements to check if the number is positive, negative, or zero
	if number > 0 {
		fmt.Println("The number is positive.")
	} else if number < 0 {
		fmt.Println("The number is negative.")
	} else {
		fmt.Println("The number is zero.")
	}
}

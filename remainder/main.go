package main

import "fmt"

func main() {
	//Declare variables
	var smallNumber, largeNumber, remainder int
	//Ask the user for a small number
	fmt.Print("Enter a small number:")
	fmt.Scan(&smallNumber)
	//Ask the user for a large number
	fmt.Print("Enter a large number:")
	fmt.Scan(&largeNumber)

	//Verify that the first number is smaller than the second number
	for {
		if largeNumber < smallNumber {
			fmt.Print("Invalid number, enter a large number:")
			fmt.Scan(&largeNumber)
			continue
		} else {
			break
		}
	}
	//Evaluate the remainder and assign it to a variable
	remainder = largeNumber % smallNumber

	//Print the Result
	fmt.Printf("The remainder of %d divided by %d is: %d\n", largeNumber, smallNumber, remainder)

}

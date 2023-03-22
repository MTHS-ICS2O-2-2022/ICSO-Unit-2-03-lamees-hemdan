// Created By: Lamees Hemdan
// Created On: March 2023
//
// This program show address

package main

import "fmt"

func main() {
	
	// variables
	var streetName string
	var streetNumber int
	
	// input
	fmt.Println("This Program gets street name and number")
	fmt.Println()
	fmt.Println("Enter street name:")
	fmt.Scanln(&streetName)
	fmt.Println("Enter street number:")
	fmt.Scanln(&streetNumber)

	// output
	fmt.Println("Your address is:", streetNumber, streetName)

	fmt.Println("\nDone.")
}

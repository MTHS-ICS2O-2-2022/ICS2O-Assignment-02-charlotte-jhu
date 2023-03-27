// Created by: Charlotte Jhu
// Created on: March 2023
//
// This program calculates the circumference of a circle

package main

import "fmt"

func main() {
	// This function calculates the circumference of a circle
	// variables
	var radius float32
	var π float32 = 3.14

	// input
	fmt.Println("This program calculates the circumference of a circle")
	fmt.Println()
	fmt.Print("Enter the radius (cm): ")
	fmt.Scanln(&radius)

	// process
	circumference := 2 * π * radius

	// output
	fmt.Println()
	fmt.Println("The circumference is", circumference, "cm")
	fmt.Println("\nDone.")
}

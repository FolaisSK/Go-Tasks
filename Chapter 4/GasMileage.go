package main

import "fmt"

func main() {
	var miles int = 0
	var gallons int = 0
	var mpg float64 = 0
	var totalMilesPerGallon float64 = 0

	var choice = 0

	for i := 0; choice != -1; i++ {
		fmt.Println("Enter Miles Driven: ")
		fmt.Scan(&miles)

		fmt.Println("Enter Gallon Used: ")
		fmt.Scan(&gallons)

		mpg = float64(miles) / float64(gallons)
		fmt.Println("MPG: ", mpg)
		totalMilesPerGallon += mpg

		fmt.Println("Enter -1 to stop operation")
		fmt.Scan(&choice)
	}

	fmt.Printf("Total MPG: %.2f\n", totalMilesPerGallon)
	fmt.Println("Thank you!")
}

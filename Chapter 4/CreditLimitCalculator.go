package main

import "fmt"

func main() {
	var (
		accountNumber    string = ""
		beginningBalance int    = 0
		charges          int    = 0
		credits          int    = 0
		creditLimit      int    = 0
		newBalance       int    = 0
	)

	fmt.Println("Enter your account number: ")
	fmt.Scan(&accountNumber)
	fmt.Println("Enter your beginning balance: ")
	fmt.Scan(&beginningBalance)
	fmt.Println("Enter your charges amount: ")
	fmt.Scan(&charges)
	fmt.Println("Enter your credits: ")
	fmt.Scan(&credits)
	fmt.Println("Enter your credits limit: ")
	fmt.Scan(&creditLimit)

	newBalance = beginningBalance + credits - charges
	fmt.Printf("New Balance is %d\n", newBalance)

	if newBalance > creditLimit {
		fmt.Println("Credit limit exceeded")
	}
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func prompt(question string) string {
	fmt.Print(question)
	input, _ := reader.ReadString('\n')
	input = strings.TrimRight(input, "\r\n")
	return input
}

func main() {
	checkout()
}

func checkout() {
	var products []string
	var quantities []float64
	var prices []float64
	var amountPaid float64 = 0

	customerName := prompt("What is the Customer's name? ")
	userResponse := ""

	for userResponse != "no" {
		product := prompt("What is the Product Name? ")
		products = append(products, product)

		quantityStr := prompt("How many pieces would you like to buy? ")
		quantity, _ := strconv.ParseFloat(quantityStr, 64)
		quantities = append(quantities, quantity)

		priceStr := prompt("How much per unit? ")
		price, _ := strconv.ParseFloat(priceStr, 64)
		prices = append(prices, price)

		userResponse = strings.ToLower(prompt("Add more Items(yes/no)? "))
	}

	totals := productTotal(quantities, prices)

	cashierName := prompt("What is your name? ")
	discountStr := prompt("How much discount will the customer get? ")
	discount, _ := strconv.ParseFloat(discountStr, 64)

	fmt.Println()

	generateInvoice(cashierName, customerName, products, quantities, prices, totals, discount)

	subTotal := calculateSubTotal(totals)
	discountAmount := calculateDiscount(subTotal, discount)
	vatAmount := calculateVAT(subTotal)
	bill := getBillTotal(subTotal, discountAmount, vatAmount)

	for amountPaid < bill {
		amountPaidStr := prompt("How much did the customer give to you? ")
		amountPaid, _ = strconv.ParseFloat(amountPaidStr, 64)
		if amountPaid < bill {
			fmt.Println("Insufficient Funds, please pay again")
		}
	}

	getReceipt(cashierName, customerName, products, quantities, prices, totals, discount, bill, amountPaid)
}

func generateInvoice(cashierName, customerName string, products []string, quantities, prices, totals []float64, discount float64) {
	fmt.Println("Generating Invoice...")
	fmt.Println("SEMICOLON STORES\nMAIN BRANCH")
	fmt.Println("LOCATION: 312, HERBERT MACAULAY WAY, SABO YABA, LAGOS.")
	fmt.Println("TEL: 090123456789\nDATE: 27-02-2026")
	fmt.Println("Cashier: ", cashierName)
	fmt.Println("Customer Name: ", customerName)
	line()
	fmt.Println("ITEM\tQTY\tPRICE\tTOTAL(NGN)")
	dash()

	for i := 0; i < len(products); i++ {
		fmt.Printf("%s\t%.2f\t%.2f\t%.2f\n", products[i], quantities[i], prices[i], totals[i])
	}

	dash()

	subTotal := calculateSubTotal(totals)
	discountTotal := calculateDiscount(subTotal, discount)
	vatTotal := calculateVAT(subTotal)

	fmt.Printf("Sub Total:         %.2f\n", subTotal)
	fmt.Printf("Discount:          %.2f\n", discountTotal)
	fmt.Printf("VAT @7.50%%:        %.2f\n", vatTotal)
	line()
	fmt.Printf("Bill Total:        %.2f\n", getBillTotal(subTotal, discountTotal, vatTotal))
	line()
	fmt.Printf("THIS IS NOT A RECEIPT KINDLY PAY %.2f\n", getBillTotal(subTotal, discountTotal, vatTotal))
	line()
}

func getReceipt(cashierName, customerName string, products []string, quantities, prices, totals []float64, discount, billTotal, amountPaid float64) {
	fmt.Println("SEMICOLON STORES\nMAIN BRANCH")
	fmt.Println("LOCATION: 312, HERBERT MACAULAY WAY, SABO YABA, LAGOS.")
	fmt.Println("TEL: 090123456789\nDATE: 27-02-2026")
	fmt.Println("Cashier: ", cashierName)
	fmt.Println("Customer Name: ", customerName)
	line()
	fmt.Println("ITEM\tQTY\tPRICE\tTOTAL(NGN)")
	dash()

	for i := 0; i < len(products); i++ {
		fmt.Printf("%s\t%.2f\t%.2f\t%.2f\n", products[i], quantities[i], prices[i], totals[i])
	}

	dash()

	subTotal := calculateSubTotal(totals)
	discountTotal := calculateDiscount(subTotal, discount)
	vatTotal := calculateVAT(subTotal)

	fmt.Printf("Sub Total:         %.2f\n", subTotal)
	fmt.Printf("Discount:          %.2f\n", discountTotal)
	fmt.Printf("VAT @7.50%%:        %.2f\n", vatTotal)
	line()
	fmt.Printf("Bill Total:        %.2f\n", getBillTotal(subTotal, discountTotal, vatTotal))
	fmt.Printf("Amount Paid:       %.2f\n", amountPaid)
	fmt.Printf("Balance:           %.2f\n", getBalance(getBillTotal(subTotal, discountTotal, vatTotal), amountPaid))
	line()
	fmt.Println("   THANK YOU FOR YOUR PATRONAGE!!!!    ")
	line()
}

func productTotal(productQuantities, productPrices []float64) []float64 {
	var totals []float64
	for i := 0; i < len(productQuantities); i++ {
		total := productPrices[i] * productQuantities[i]
		totals = append(totals, total)
	}
	return totals
}

func calculateSubTotal(totals []float64) float64 {
	subTotal := 0.0
	for _, total := range totals {
		subTotal += total
	}
	return subTotal
}

func calculateDiscount(subTotal, discount float64) float64 {
	discount = discount / 100
	return subTotal * discount
}

func calculateVAT(subTotal float64) float64 {
	vatAmount := 0.075
	return subTotal * vatAmount
}

func getBillTotal(subTotal, discount, vatAmount float64) float64 {
	bill := subTotal + vatAmount - discount
	return bill
}

func getBalance(billTotal, amountPaid float64) float64 {
	return amountPaid - billTotal
}

func line() {
	fmt.Println("=======================================================")
}

func dash() {
	fmt.Println("--------------------------------------------------------")
}

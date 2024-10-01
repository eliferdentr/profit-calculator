package main

import "fmt"

func main() {
	var revenue, expenses, taxRate float64

	fmt.Print("Please enter revenue: ")
	fmt.Scan(&revenue)
	fmt.Print("\nPlease enter expenses: ")
	fmt.Scan(&expenses)
	fmt.Print("\nPlease enter tax rate (in decimal): ")
	fmt.Scan(&taxRate)

	ebt := calculateEbt(revenue, expenses)
	eat := calculateEat(revenue, expenses, taxRate)

	fmt.Printf("Earning before taxes: %.2f\n", ebt)
	fmt.Printf("\nEarning after taxes: %.2f\n ", eat)

	ratio := calculateRatio(ebt, eat)
	if ratio != -1 {
		fmt.Print("\nRatio: ", calculateRatio(ebt, eat))
	} else {
		fmt.Print("ratio can't be calculated because the profit is 0")
	}

}

func calculateEbt(revenue float64, expenses float64) float64 {
	return revenue - expenses
}

func calculateEat(revenue float64, expenses float64, taxRate float64) float64 {
	var netIncome = calculateEbt(revenue, expenses)
	var tax = netIncome * taxRate
	return netIncome - tax
}

func calculateRatio(ebt float64, profit float64) float64 {
	return ebt / profit
}

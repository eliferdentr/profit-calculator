package main

import "fmt"

func main() {
	revenue := takeUserInput("Please enter revenue: ")
	expenses := takeUserInput("Please enter expenses: ")
	taxRate :=  takeUserInput("Please enter tax rate (in decimal): ")

	ebt := calculateEbt(revenue, expenses)
	eat := calculateEat(revenue, expenses, taxRate)

	// eat, ebt := calculateEatEbtRatio(revenue, expenses, taxRate)

	formattedEbt := fmt.Sprintf("Earning before taxes: %.2f\n", ebt)
	formattedEat := fmt.Sprintf("Earning after taxes: %.2f\n ", eat)
	fmt.Print(formattedEbt, formattedEat)

	ratio := calculateRatio(ebt, eat)

	if ratio != -1 {
		fmt.Printf("\nRatio: %.2f\n", ratio)
	} else {
		fmt.Printf("ratio can't be calculated because the profit is 0")
	}

}

func calculateEbt(revenue, expenses float64) float64 {
	return revenue - expenses
}

func calculateEat(revenue, expenses, taxRate float64) float64 {
	netIncome := calculateEbt(revenue, expenses)
	tax := netIncome * taxRate
	return netIncome - tax
}

func calculateRatio(ebt, profit float64) float64 {
	return ebt / profit
}

func takeUserInput (outputText string) (inputValue float64) {
	fmt.Println(outputText)
	fmt.Scan(&inputValue)
	return
}

// func calculateEatEbtRatio (revenue, expenses, taxRate float64) (float64, float64){
// 	ebt := revenue - expenses
// 	eat := ebt - ebt * taxRate
// 	return ebt, eat
// }

// func calculateEatEbtRatio (revenue, expenses, taxRate float64) (ebt float64, eat float64){
// 	ebt = revenue - expenses
// 	eat = ebt - ebt * taxRate
// 	return ebt, eat
// }
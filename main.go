package main

import (
	"errors"
	"fmt"
	"os"
)

/*
validate user input: not negative or 0. if so, generate an error and handle it.
show error message and exit if input is invalid
store calculated

store the calculated values in a file.
*/

const logFile = "log.txt"
const ebtLogString  = "EBT is: "
const eatLogString = "EAT is: "
const ratioLogString  = "Ratio is: "

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
	ebt := revenue - expenses
	
	logOperation(ebtLogString, ebt)
	return ebt
}

func calculateEat(revenue, expenses, taxRate float64) float64 {
	netIncome := calculateEbt(revenue, expenses)
	tax := netIncome * taxRate
	eat := netIncome - tax
	logOperation(eatLogString, eat)
	return eat
}

func calculateRatio(ebt, profit float64) float64 {
	ratio := ebt / profit
	logOperation(ratioLogString, ratio)
	return ratio
}

func takeUserInput (outputText string) (inputValue float64) {
	fmt.Println(outputText)
	fmt.Scan(&inputValue)
	if err := validateUserInput(inputValue); err != nil {
		fmt.Println("Error! : ", err)
		os.Exit(1)
		return
	}
	return
}

func validateUserInput(userInput float64) error{
	if userInput <= 0 {
		return errors.New("input cannot be less than or equal to 0")
	}
	return nil
}

func logOperation (operation string, value float64) {
	text := fmt.Sprintf("%s %.2f \n", operation, value)
	file, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		os.Exit(1)
	}
	defer file.Close()
	if _, err := file.WriteString(text); err != nil{
		fmt.Println("Error writing to the file: ",err)
		os.Exit(1)
	}}

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
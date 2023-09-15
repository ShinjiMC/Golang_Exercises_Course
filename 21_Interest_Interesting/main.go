package main

import (
	"fmt"
)

// InterestRate returns the interest rate for the provided balance.
func InterestRate(balance float64) float32 {
	if balance < 0.0 {
		return 3.213
	}
	if 1000.0 > balance {
		return 0.5
	}
	if 5000.0 > balance {
		return 1.621
	}
	return 2.475
}

// Interest calculates the interest for the provided balance.
func Interest(balance float64) float64 {
	return balance * (float64(InterestRate(balance)) / 100.0)
}

// AnnualBalanceUpdate calculates the annual balance update, taking into account the interest rate.
func AnnualBalanceUpdate(balance float64) float64 {
	return balance + Interest(balance)
}

// YearsBeforeDesiredBalance calculates the minimum number of years required to reach the desired balance.
func YearsBeforeDesiredBalance(balance, targetBalance float64) int {
	var accumulatingBalance = balance
	years := 0
	for accumulatingBalance < targetBalance {
		accumulatingBalance = AnnualBalanceUpdate(accumulatingBalance)
		years++
	}
	return years
}

func main() {
	balance := 200.75
	fmt.Println("Test 1: ", InterestRate(balance))
	fmt.Println("Test 2: ", Interest(balance))
	fmt.Println("Test 3: ", AnnualBalanceUpdate(balance))
	targetBalance := 214.88
	fmt.Println("Test 4: ", YearsBeforeDesiredBalance(balance, targetBalance))
}

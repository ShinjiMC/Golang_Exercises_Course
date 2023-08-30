package main

import (
	"fmt"
	"strings"
)

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return "Welcome to the Tech Palace, " + strings.ToUpper(customer)
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	var pointers = strings.Repeat("*", numStarsPerLine)
	return pointers + "\n" + welcomeMsg + "\n" + pointers
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	return strings.TrimSpace(strings.ReplaceAll(oldMsg, "*", ""))
}

func main() {
	message := `
	**************************
	*    BUY NOW, SAVE 10%   *
	**************************
	`
	fmt.Println("Test 1: ", WelcomeMessage("Judy"))
	fmt.Println("Test 2:")
	fmt.Println(AddBorder("Welcome!", 10))
	fmt.Println("Test 3:")
	fmt.Println(CleanupMessage(message))
}

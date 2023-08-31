package main

import "fmt"

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	var ret int
	switch card {
	case "ace":
		ret = 11
	case "two":
		ret = 2
	case "three":
		ret = 3
	case "four":
		ret = 4
	case "five":
		ret = 5
	case "six":
		ret = 6
	case "seven":
		ret = 7
	case "eight":
		ret = 8
	case "nine":
		ret = 9
	case "ten", "jack", "queen", "king":
		ret = 10
	default:
		ret = 0
	}
	return ret
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	sum := ParseCard(card1) + ParseCard(card2)
	var rta string
	if sum == 22 {
		rta = "P"
	} else if sum == 21 {
		if ParseCard(dealerCard) < 10 {
			rta = "W"
		} else {
			rta = "S"
		}
	} else if sum <= 20 && sum >= 17 {
		rta = "S"
	} else if sum <= 16 && sum >= 12 {
		if ParseCard(dealerCard) >= 7 {
			rta = "H"
		} else {
			rta = "S"
		}
	} else {
		rta = "H"
	}
	return rta
}

func main() {
	fmt.Println(FirstTurn("ace", "ace", "jack"))
	fmt.Println(FirstTurn("ace", "king", "ace"))
	fmt.Println(FirstTurn("five", "queen", "ace"))
}

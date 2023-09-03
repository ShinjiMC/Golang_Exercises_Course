package main

import "fmt"

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	return map[string]int{
		"quarter_of_a_dozen": 3,
		"half_of_a_dozen":    6,
		"dozen":              12,
		"small_gross":        120,
		"gross":              144,
		"great_gross":        1728,
	}
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return make(map[string]int)
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	if _, validate := units[unit]; !validate {
		return false
	}
	bill[item] += units[unit]
	return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	_, unitInUnits := units[unit]
	_, itemInBill := bill[item]
	if !unitInUnits || !itemInBill || bill[item]-units[unit] < 0 {
		return false
	}
	bill[item] -= units[unit]
	if bill[item] == 0 {
		delete(bill, item)
	}
	return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	units, itemInBill := bill[item]
	if !itemInBill {
		return 0, false
	}
	return units, true
}

func main() {
	bill := NewBill()
	units := Units()
	ok := RemoveItem(bill, units, "carrot", "dozen")
	fmt.Println(ok)
	bill = map[string]int{"carrot": 12, "grapes": 3}
	qty, ok := GetItem(bill, "carrot")
	fmt.Println(qty)
	fmt.Println(ok)
}

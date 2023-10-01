package main

import (
	"fmt"
)

// Record represents an expense record.
type Record struct {
	Day      int
	Amount   float64
	Category string
}

// DaysPeriod represents a period of days for expenses.
type DaysPeriod struct {
	From int
	To   int
}

func (p DaysPeriod) Compare(d int) bool {
	return p.From <= d && d <= p.To
}

// Filter returns the records for which the predicate function returns true.
func Filter(in []Record, predicate func(Record) bool) []Record {
	var res []Record
	for _, val := range in {
		if predicate(val) {
			res = append(res, val)
		}
	}
	return res
}

// ByDaysPeriod returns predicate function that returns true when
// the day of the record is inside the period of day and false otherwise.
func ByDaysPeriod(p DaysPeriod) func(Record) bool {
	return func(r Record) bool {
		return p.Compare(r.Day)
	}
}

// ByCategory returns predicate function that returns true when
// the category of the record is the same as the provided category
// and false otherwise.
func ByCategory(c string) func(Record) bool {
	return func(r Record) bool {
		return c == r.Category
	}
}

// TotalByPeriod returns total amount of expenses for records
// inside the period p.
func TotalByPeriod(in []Record, p DaysPeriod) float64 {
	res := Filter(in, ByDaysPeriod(p))
	var asn float64
	for _, val := range res {
		asn += val.Amount
	}
	return asn
}

// CategoryExpenses returns total amount of expenses for records
// in category c that are also inside the period p.
// An error must be returned only if there are no records in the list that belong
// to the given category, regardless of period of time.
func CategoryExpenses(in []Record, p DaysPeriod, c string) (float64, error) {
	res := Filter(in, ByCategory(c))
	if len(res) == 0 {
		return 0, fmt.Errorf("unknown category %s", c)
	}
	return TotalByPeriod(res, p), nil
}

func Day1Records(r Record) bool {
	return r.Day == 1
}

func main() {
	records := []Record{
		{Day: 1, Amount: 15, Category: "groceries"},
		{Day: 11, Amount: 300, Category: "utility-bills"},
		{Day: 12, Amount: 28, Category: "groceries"},
		{Day: 26, Amount: 300, Category: "university"},
		{Day: 28, Amount: 1300, Category: "rent"},
	}
	period := DaysPeriod{From: 1, To: 15}
	p1 := DaysPeriod{From: 1, To: 15}
	fmt.Println("Test 1: ", Filter(records, Day1Records))
	fmt.Println("Test 2: ", Filter(records, ByDaysPeriod(period)))
	fmt.Println("Test 3: ", Filter(records, ByCategory("groceries")))
	fmt.Println("Test 4: ", TotalByPeriod(records, p1))
	v, r := CategoryExpenses(records, p1, "groceries")
	fmt.Printf("Test 5:  %f - %v\n", v, r)
}

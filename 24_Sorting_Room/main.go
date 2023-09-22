package main

import (
	"fmt"
	"strconv"
)

const numberFormat = "This is%s the number %.1f"

// DescribeNumber should return a string describing the number.
func DescribeNumber(f float64) string {
	return fmt.Sprintf(numberFormat, "", f)
}

type NumberBox interface {
	Number() int
}

type numberBoxContaining struct {
	n int
}

func (nb numberBoxContaining) Number() int {
	return nb.n
}

// DescribeNumberBox should return a string describing the NumberBox.
func DescribeNumberBox(nb NumberBox) string {
	return fmt.Sprintf(numberFormat, " a box containing", float64(nb.Number()))
}

type FancyNumber struct {
	n string
}

func (i FancyNumber) Value() string {
	return i.n
}

type FancyNumberBox interface {
	Value() string
}

// ExtractFancyNumber should return the integer value for a FancyNumber
// and 0 if any other FancyNumberBox is supplied.
func ExtractFancyNumber(fnb FancyNumberBox) int {
	fancyNum, ok := fnb.(FancyNumber)
	if !ok {
		return 0
	}
	num, _ := strconv.Atoi(fancyNum.Value())
	return num
}

// DescribeFancyNumberBox should return a string describing the FancyNumberBox.
func DescribeFancyNumberBox(fnb FancyNumberBox) string {
	return fmt.Sprintf(numberFormat, " a fancy box containing", float64(ExtractFancyNumber(fnb)))
}

// DescribeAnything should return a string describing whatever it contains.
func DescribeAnything(i interface{}) string {
	switch v := i.(type) {
	case int:
		return DescribeNumber(float64(v))
	case float64:
		return DescribeNumber(v)
	case NumberBox:
		return DescribeNumberBox(v)
	case FancyNumberBox:
		return DescribeFancyNumberBox(v)
	default:
		return "Return to sender"
	}
}

func main() {
	fmt.Println("Test 1: ", DescribeAnything(-12.345))
	fmt.Println("Test 2: ", DescribeAnything(42))
	fmt.Println("Test 3: ", DescribeAnything(numberBoxContaining{12}))
	fmt.Println("Test 4: ", DescribeAnything(FancyNumber{"1256"}))
	fmt.Println("Test 5: ", DescribeAnything("some string"))
}

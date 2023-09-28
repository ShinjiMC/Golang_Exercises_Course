package main

import (
	"errors"
	"fmt"
)

type testFodderCalculator struct {
	amount    float64
	amountErr error
	factor    float64
	factorErr error
}

func (fc testFodderCalculator) FodderAmount(int) (float64, error) {
	return fc.amount, fc.amountErr
}

func (fc testFodderCalculator) FatteningFactor() (float64, error) {
	return fc.factor, fc.factorErr
}

type FodderCalculator interface {
	FodderAmount(int) (float64, error)
	FatteningFactor() (float64, error)
}

// TODO: define the 'DivideFood' function
func DivideFood(unit FodderCalculator, cows int) (float64, error) {
	fodder, err := unit.FodderAmount(cows)
	if err != nil {
		return 0, err
	}
	multiply, err2 := unit.FatteningFactor()
	if err2 != nil {
		return 0, err2
	}
	return fodder / float64(cows) * multiply, nil
}

// TODO: define the 'ValidateInputAndDivideFood' function
func ValidateInputAndDivideFood(unit FodderCalculator, cows int) (float64, error) {
	if 0 < cows {
		return DivideFood(unit, cows)
	}
	return 0, errors.New("invalid number of cows")
}

type InvalidCowsError struct {
	cows    int
	message string
}

func (e *InvalidCowsError) Error() string {
	return fmt.Sprintf("%d cows are invalid: %s", e.cows, e.message)
}

// TODO: define the 'ValidateNumberOfCows' function
func ValidateNumberOfCows(cows int) error {
	if cows < 0 {
		return &InvalidCowsError{
			cows:    cows,
			message: "there are no negative cows",
		}
	}
	if cows == 0 {
		return &InvalidCowsError{
			cows:    cows,
			message: "no cows don't need food",
		}
	}
	return nil
}

func main() {
	cowUnit := testFodderCalculator{
		amount: 100,
		factor: 1,
	}
	cows := 10
	foodPerCow, err := DivideFood(cowUnit, cows)
	if err != nil {
		fmt.Printf("Error dividing food: %.2f - %v\n", foodPerCow, err)
	} else {
		fmt.Printf("Food per cow: %.2f - %v\n", foodPerCow, err)
	}
	foodPerValidCow, err := ValidateInputAndDivideFood(cowUnit, 5)
	if err != nil {
		fmt.Printf("Error validating and dividing food: %.2f - %v\n", foodPerValidCow, err)
	} else {
		fmt.Printf("Food per cow (valid): %.2f - %v\n", foodPerValidCow, err)
	}
	foodPerValidCow, err = ValidateInputAndDivideFood(cowUnit, -1)
	if err != nil {
		fmt.Printf("Error validating and dividing food (invalid): %.2f - %v\n", foodPerValidCow, err)
	}
	fmt.Printf("Error validating the number of cows (valid): %v\n", ValidateNumberOfCows(5))
	fmt.Printf("Error validating the number of cows (invalid): %v\n", ValidateNumberOfCows(-1))
	fmt.Printf("Error validating the number of cows (zero): %v\n", ValidateNumberOfCows(0))
}

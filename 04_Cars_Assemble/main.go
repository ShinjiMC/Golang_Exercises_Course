package main

import "fmt"

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return float64(productionRate) * float64(successRate/100)
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	return int(float64(float64(productionRate/60) * float64(successRate/100)))
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	var unitCars = carsCount % 10
	var tenCars = carsCount / 10
	return uint((tenCars * 95000) + (unitCars * 10000))
}

func main() {
	fmt.Println("Workings Cars per Hour: ", CalculateWorkingCarsPerHour(1547, 90))
	fmt.Println("Workings Cars per Minute: ", CalculateWorkingCarsPerMinute(1105, 90))
	fmt.Println("Total Cost of Cars", CalculateCost(37))
}

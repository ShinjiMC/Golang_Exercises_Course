package main

import "fmt"

const OvenTime = 40

func RemainingOvenTime(actual int) int {
	return OvenTime - actual
}

func PreparationTime(numberOfLayers int) int {
	return numberOfLayers * 2
}

func ElapsedTime(numberOfLayers, actualMinutesInOven int) int {
	return PreparationTime(numberOfLayers) + actualMinutesInOven
}

func main() {
	fmt.Println("Tiempo Transcurrido: ", ElapsedTime(3, 20))
	fmt.Println("Tiempo Restante: ", RemainingOvenTime(ElapsedTime(3, 20)))
}

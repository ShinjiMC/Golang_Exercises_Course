package main

import "fmt"

func CanFastAttack(knightIsAwake bool) bool {
	return !knightIsAwake
}
func CanSpy(knightIsAwake, archerIsAwake, prisonerIsAwake bool) bool {
	return (knightIsAwake || archerIsAwake || prisonerIsAwake)
}

func CanSignalPrisoner(archerIsAwake, prisonerIsAwake bool) bool {
	return (!archerIsAwake && prisonerIsAwake)
}

func CanFreePrisoner(knightIsAwake, archerIsAwake, prisonerIsAwake, petDogIsPresent bool) bool {
	if (!archerIsAwake && petDogIsPresent) || (!petDogIsPresent && prisonerIsAwake && !archerIsAwake && !knightIsAwake) {
		return true
	}
	return false
}

func main() {
	var knightIsAwake = false
	var archerIsAwake = false
	var prisonerIsAwake = true
	var petDogIsPresent = true
	fmt.Println("Puede Espiar: ", CanSpy(knightIsAwake, archerIsAwake, prisonerIsAwake))
	fmt.Println("Puede Atacar Rapido: ", (CanFastAttack(knightIsAwake)))
	fmt.Println("Puede Senalar: ", CanSignalPrisoner(archerIsAwake, prisonerIsAwake))
	fmt.Println("Puede Rescatar: ", CanFreePrisoner(knightIsAwake, archerIsAwake, prisonerIsAwake, petDogIsPresent))
}

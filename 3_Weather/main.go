// Package main: show current weather of a current location
package main

import "fmt"

var CurrentCondition string //CurrentCondition is a current weather of a location.
var CurrentLocation string  //CurrentLocation is a current city to find weather.

func Forecast(city, condition string) string { //Forecast Func shows the current weather of a city.
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}

func main() {
	CurrentCondition = "Sunny"
	CurrentLocation = "Arequipa"
	fmt.Println(Forecast(CurrentCondition, CurrentLocation))
}

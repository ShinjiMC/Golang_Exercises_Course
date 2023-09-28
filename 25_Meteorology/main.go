package main

import (
	"fmt"
)

type TemperatureUnit int

const (
	Celsius    TemperatureUnit = 0
	Fahrenheit TemperatureUnit = 1
)

// Add a String method to the TemperatureUnit type
func (unit TemperatureUnit) String() string {
	return [...]string{"°C", "°F"}[unit]
}

type Temperature struct {
	degree int
	unit   TemperatureUnit
}

// Add a String method to the Temperature type
func (unit Temperature) String() string {
	return fmt.Sprintf("%d %s", unit.degree, unit.unit)
}

type SpeedUnit int

const (
	KmPerHour    SpeedUnit = 0
	MilesPerHour SpeedUnit = 1
)

// Add a String method to SpeedUnit
func (unit SpeedUnit) String() string {
	return [...]string{"km/h", "mph"}[unit]
}

type Speed struct {
	magnitude int
	unit      SpeedUnit
}

// Add a String method to Speed
func (unit Speed) String() string {
	return fmt.Sprintf("%d %s", unit.magnitude, unit.unit)
}

type MeteorologyData struct {
	location      string
	temperature   Temperature
	windDirection string
	windSpeed     Speed
	humidity      int
}

// Add a String method to MeteorologyData
func (unit MeteorologyData) String() string {
	return fmt.Sprintf("%s: %s, Wind %s at %s, %d%% Humidity", unit.location, unit.temperature, unit.windDirection, unit.windSpeed, unit.humidity)
}

func main() {
	celsiusUnit := Celsius
	fmt.Println("Test 1: ", celsiusUnit.String())
	fahrenheitTemp := Temperature{
		degree: 75,
		unit:   Fahrenheit,
	}
	fmt.Println("Test 2: ", fmt.Sprint(fahrenheitTemp))
	mphUnit := MilesPerHour
	fmt.Println("Test 3: ", mphUnit.String())
	windSpeedNow := Speed{
		magnitude: 18,
		unit:      KmPerHour,
	}
	fmt.Println("Test 4: ", windSpeedNow.String())
	sfData := MeteorologyData{
		location: "San Francisco",
		temperature: Temperature{
			degree: 57,
			unit:   Fahrenheit,
		},
		windDirection: "NW",
		windSpeed: Speed{
			magnitude: 19,
			unit:      MilesPerHour,
		},
		humidity: 60,
	}
	fmt.Println("Test 5: ", sfData.String())
}

package main

import "fmt"

// TODO: define the 'Car' type struct
type Car struct {
	speed        int
	batteryDrain int
	battery      int
	distance     int
}

// NewCar creates a new remote controlled car with full battery and given specifications.
func NewCar(speed, batteryDrain int) Car {
	return Car{speed: speed, batteryDrain: batteryDrain, battery: 100, distance: 0}
}

// TODO: define the 'Track' type struct
type Track struct {
	distance int
}

// NewTrack creates a new track
func NewTrack(distance int) Track {
	return Track{distance: distance}
}

// TODO: define the 'Drive()' method
func (c *Car) Drive() {
	if c.battery >= c.batteryDrain {
		c.battery -= c.batteryDrain
		c.distance += c.speed
	}
}

// TODO: define the 'DisplayDistance() string' method
func (c Car) DisplayDistance() string {
	return fmt.Sprintf("Driven %d meters", c.distance)
}

// TODO: define the 'DisplayBattery() string' method
func (c Car) DisplayBattery() string {
	return fmt.Sprintf("Battery at %d%%", c.battery)
}

// TODO: define the 'CanFinish(trackDistance int) bool' method
func (c Car) CanFinish(trackdistance int) bool {
	for c.battery >= c.batteryDrain {
		c.Drive()
		if c.distance > trackdistance {
			break
		}
	}
	return c.distance >= trackdistance
}

func main() {
	speed := 5
	batteryDrain := 2
	car := NewCar(speed, batteryDrain)
	distance := 100
	track := NewTrack(distance)
	fmt.Println(car.DisplayDistance())
	fmt.Println(car.DisplayBattery())
	fmt.Println(car.CanFinish(track.distance))
}

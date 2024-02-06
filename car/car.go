package car

import "fmt"

// Declares the Car struct to house the concept of a car
type Car struct {
	Name       string
	Doors      int
	CoolNumber int
}

func (c Car) String() string {
	return fmt.Sprintf("Car: {Name: %s, Doors: %d}", c.Name, c.Doors)
}

// This function is the designated initializer for the Car
func NewCar(carname string, cardoors int) Car {
	car := Car{Name: carname, Doors: cardoors}
	car.CoolNumber = 4 * car.Doors
	return car
}

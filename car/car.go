package car

import "fmt"

type Car struct {
	Name       string
	Doors      int
	CoolNumber int
}

func (c Car) String() string {
	return fmt.Sprintf("Car: {Name: %s, Doors: %d}", c.Name, c.Doors)
}

func NewCar(carname string, cardoors int) Car {
	car := Car{Name: carname, Doors: cardoors}
	car.CoolNumber = 4 * car.Doors
	return car
}
package method

import "fmt"

// Engine toy engine
type Engine interface {
	Start()
	Stop()
}

// Car include Engine
type Car struct {
	Engine
	wheelCount int
}

// SuperCar is a car
type SuperCar struct {
	Car
}

// GoToWorkIn Start a Car
func (c *Car) GoToWorkIn() {
	c.Start()
	c.Stop()
}

func (c *Car) Start() {
	fmt.Printf("Here %v started\n", c)
}
func (c *Car) Stop() {
	fmt.Printf("Here %T stoped\n", c)
}
func (c *Car) numberOfWheels() int {
	return c.wheelCount
}

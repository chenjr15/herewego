package itfsandrflx

import "fmt"

// Car A simple Car type
type Car struct {
	name  string
	speed int
}

func (c *Car) run() {
	fmt.Printf("%s is running with speed %d. ", c.name, c.speed)
}

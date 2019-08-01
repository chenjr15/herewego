package method

import "testing"

func TestCar(t *testing.T) {
	car := Car{wheelCount: 4}
	superCar := SuperCar{Car: Car{wheelCount: 100}}
	car.GoToWorkIn()
	wheelCount := car.numberOfWheels()
	t.Logf("Car have %d wheels.", wheelCount)
	superCar.GoToWorkIn()
	wheelCount = superCar.numberOfWheels()
	t.Logf("SuperCar have %d wheels.", wheelCount)

}

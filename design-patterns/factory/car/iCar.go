package car

type iCar interface {
	setName(name string)
	getName() string
	setPower(power int)
	getPower() int
}

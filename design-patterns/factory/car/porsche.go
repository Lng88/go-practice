package car

type porsche struct {
	car
}

func NewPorsche(name string, power int) iCar {
	return &porsche{
		car: car{
			name:  name,
			power: power,
		},
	}
}

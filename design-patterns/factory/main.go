package main

import (
	"fmt"

	"car.example/car"
)

func main() {
	turbo := car.NewPorsche("Porsche Turbo", 10)
	nineEleven := car.NewPorsche("Porsche 911", 8)
	fmt.Println(turbo, nineEleven)
}

package main

import "fmt"

type Vehicle struct {
	NumberOfWheels     int
	NumberofPassengers int
}

type Car struct {
	Make      string
	Model     string
	Year      int
	IsElectic bool
	IsHybrid  bool
	Vehicle   Vehicle
}

func (v Vehicle) showDetails() {
	fmt.Println("Number of passengers:", v.NumberofPassengers)
	fmt.Println("Number of wheels:", v.NumberOfWheels)
}

func (c Car) show() {
	fmt.Println("Make:", c.Make)
	fmt.Println("Model:", c.Model)
	fmt.Println("Year:", c.Year)
	fmt.Println("Is electric:", c.IsElectic)
	fmt.Println("Is hybrid:", c.IsHybrid)
	c.Vehicle.showDetails()
}

func main() {
	suv := Vehicle{
		NumberOfWheels:     4,
		NumberofPassengers: 6,
	}
	volvoXC90 := Car{
		Make:      "Volvo",
		Model:     "XC90 T8",
		Year:      2021,
		IsElectic: false,
		IsHybrid:  true,
		Vehicle:   suv,
	}
	volvoXC90.show()
}

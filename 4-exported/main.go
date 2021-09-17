package main

import "myapp/staff"

var employees = []staff.Employee{
	{FirstName: "John", LastName: "Smith", Salary: 30000, FullTime: true},
	{FirstName: "Sally", LastName: "Jones", Salary: 50000, FullTime: true},
	{FirstName: "Eric", LastName: "Jones", Salary: 14000, FullTime: false},
}

func main() {
	myStaff := staff.Office{
		AllStaff: employees,
	}
}

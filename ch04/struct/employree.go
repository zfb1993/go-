package main

import (
	"time"
	"fmt"
)

type Employee struct{
	ID int
	Name string
	Address string
	DoB time.Time
	Position string
	Salary int
}

func main(){
	var dilbert Employee
	var employeeOfTheMonth *Employee = &dilbert
	employeeOfTheMonth.Position += " (proactive team player)"
	fmt.Println(*employeeOfTheMonth)
}
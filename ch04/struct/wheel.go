package main

import "fmt"

type Point struct{
	X,Y int
}

type Circle struct {
	Point
	Radius int
}

type Wheel struct {
	Circle
	Spokes int
}

func main(){
	var w Wheel
	w = Wheel{
		Circle:Circle{
			Point: Point{X: 8, Y: 8},
			Radius:20,
		},
		Spokes:5,
	}
	fmt.Println(w)
}
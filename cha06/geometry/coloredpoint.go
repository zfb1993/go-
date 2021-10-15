package main

import (
	"fmt"
	"image/color"
)

type Point struct {
	X,Y float64
}

type ColordPoint struct {
	Point 
	Color color.RGBA
}

func main()  {
	var cp ColordPoint
	cp.X = 5
	fmt.Println(cp.Point.X)

	cp.Point.Y = 2
	fmt.Println(cp.Y) // "2"
}
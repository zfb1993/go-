package main

import (
	"math"
	"fmt"
)

type Point struct {X,Y float64}

type Path []Point

func (p Point) Distance (q Point) float64{
    return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func main(){
	p := Point{1, 2}
	q := Point{4, 6}
	fmt.Println(p.Distance(q))  // "5", method call

	perim := Path{
		{1, 1},
		{5, 1},
		{5, 4},
		{1, 1},
	}
	fmt.Println(perim.Distance()) // "12"
}

func (path Path)  Distance() float64{
	sum := 0.0
    for i := range path {
        if i > 0 {
            sum += path[i-1].Distance(path[i])
        }
    }
    return sum
}
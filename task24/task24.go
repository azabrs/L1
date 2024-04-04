package main

import (
	"fmt"
	"math"
)

type Point struct{
	x float64
	y float64
}

func NewPoint(x, y float64) Point{
	return Point{
		x : x,
		y : y,
	}
}

func (p1 Point) Distance(p2 Point) float64{
	return math.Sqrt(math.Pow((p2.x-p1.x), 2) + math.Pow((p2.y-p1.y), 2)) //calculate the distance between 2 points using formula
}

func main(){
	point1 := NewPoint(1, 3) // create first point
	point2 := NewPoint(2, 8) // create second point
	fmt.Println(point1.Distance(point2))
	
}
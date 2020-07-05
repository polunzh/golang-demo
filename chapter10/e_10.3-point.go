package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

type Point3 struct {
	x, y, z float32
}

type Polar struct {
	R, L float64
}

func Abs(p *Point) float64 {
	return math.Sqrt(float64(p.x*p.x + p.y*p.y))
}

func Scale(p *Point, s float64) {
	p.x *= s
	p.y *= s
	return
}

func main() {
	var p = Point{10, 10}
	fmt.Printf("%f\n", Abs(&p))
	Scale(&p, 2)
	fmt.Printf("%v", p)
}

package main

import (
	"fmt"
	"math"
)

type Point struct {
	X float64
	Y float64
}

type Vector struct {
	Start Point
	End   Point
}

func main() {
	p1 := Point{X: 1, Y: 3}
	p2 := Point{X: 10, Y: 20}
	p3 := Point{X: -3, Y: -13}
	p4 := Point{X: -30, Y: 22}

	v1 := Vector{Start: p1, End: p2}
	v2 := Vector{Start: p3, End: p4}

	l1 := length(v1)
	fmt.Println("Lenght of vector 1: ", math.Round(l1))

	l2 := length(v2)
	fmt.Println("Lenght of vector 2: ", math.Round(l2))

	resultant := add(v1, v2)
	fmt.Println("Resultant of vectors: ", resultant)

	difference := subtract(v1, v2)
	fmt.Println("Difference of vectors: ", difference)

	multiplyV1 := multiply(v1, 5)
	fmt.Println("Scale v1: ", multiplyV1)

	multiplyV2 := multiply(v2, 5)
	fmt.Println("Scale v2: ", multiplyV2)

	dotProduct := dotProduct(v1, v2)
	fmt.Println("Dot Product: ", dotProduct)

	crossProduct := crossProduct(v1, v2)
	fmt.Println("Cross Product: ", crossProduct)
}

func length(v Vector) float64 {
	return math.Sqrt((v.End.X-v.Start.X)*(v.End.X-v.Start.X) + (v.End.Y-v.Start.Y)*(v.End.Y-v.Start.Y))
}

func add(v1 Vector, v2 Vector) Vector {
	p1 := Point{X: v1.Start.X + v2.Start.X, Y: v1.Start.Y + v2.Start.Y}
	p2 := Point{X: v1.End.X + v2.End.X, Y: v1.End.Y + v2.End.Y}
	return Vector{p1, p2}
}

func subtract(v1 Vector, v2 Vector) Vector {
	p1 := Point{X: v1.Start.X - v2.Start.X, Y: v1.Start.Y - v2.Start.Y}
	p2 := Point{X: v1.End.X - v2.End.X, Y: v1.End.Y - v2.End.Y}
	return Vector{p1, p2}
}

func multiply(v Vector, scalar float64) Vector {
	p1 := v.Start
	p2 := v.End

	p1.X = p1.X * scalar
	p1.Y = p1.Y * scalar

	p2.X = p2.X * scalar
	p2.Y = p2.Y * scalar

	return Vector{p1, p2}
}

func dotProduct(v1 Vector, v2 Vector) float64 {
	ax := v1.End.X - v1.Start.X
	ay := v1.End.Y - v1.Start.Y
	bx := v2.End.X - v2.Start.X
	by := v2.End.Y - v2.Start.Y

	return ax*bx + ay*by
}

func crossProduct(v1 Vector, v2 Vector) float64 {
	return (v1.Start.X-v1.End.X)*(v2.Start.Y-v2.End.Y) - (v1.Start.Y-v1.End.Y)*(v2.Start.X-v2.End.X)
}
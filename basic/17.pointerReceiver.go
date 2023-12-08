package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

type ColouredPoint struct {
	Point
	colour string
}

func (p Point) DistanceFrom(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// pointer by reference
func (p *Point) Scale(factor float64) {
	if p == nil {
		return
	}
	p.X = p.X * factor
	p.Y = p.Y * factor
}

func main() {
	p1 := Point{1, 1}
	p2 := Point{1, 3}
	var p3 *Point
	fmt.Println(p1)
	fmt.Println(p1.DistanceFrom(p2))
	// pointer receiver
	p1.Scale(4)
	fmt.Println(p1)
	// nil pointer receiver
	p3.Scale(5)
	// method as function values
	scaleFn := p2.Scale
	scaleFn(3)
	fmt.Println(p2)
	// Composition by struct embedding
	p4 := ColouredPoint{Point{3, 6}, "green"}
	colouredScaleFn := p4.Point.Scale
	colouredScaleFn(6)
	fmt.Println(p4)
}

/**
PS D:\Study\Go\GoLearn\basic> go run 17.pointerReceiver.go
{1 1}
2
{4 4}
{3 9}
*/

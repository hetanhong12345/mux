// hello
package main

import (
	"math"
	"fmt"
)

type Point struct {
	X float64
	Y float64
}

func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)

}

func main() {
	p := Point{1, 2}
	q := Point{4, 6}

	distanceFromP := p.Distance        // method value
	fmt.Println(distanceFromP(q))      // "5"
	var origin Point                   // {0, 0}
	fmt.Println(distanceFromP(origin)) // "2.23606797749979", sqrt(5)

	distance := Point.Distance   // method expression
	fmt.Println(distance(p, q))  // "5"
	fmt.Printf("%V\n", distance) // "func(Point, Point) float64"

}

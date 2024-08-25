package utils

import "math"

type Point2 struct {
	X int
	Y int
}

func (p Point2) Add(p2 Point2) Point2 {
	return Point2{p.X + p2.X, p.Y + p2.Y}
}

func (p Point2) AddX(x int) Point2 {
	return Point2{p.X + x, p.Y}
}

func (p Point2) AddY(y int) Point2 {
	return Point2{p.X, p.Y + y}
}

func (p Point2) Sub(p2 Point2) Point2 {
	return Point2{p.X - p2.X, p.Y - p2.Y}
}

// Returns a slice of points in a 2D grid of the given width and height.
// Iterate over the points row by row, left to right, starting from 0,0.
func PointsIn2D(width, height int) []Point2 {
	points := make([]Point2, 0, width*height)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			points = append(points, Point2{x, y})
		}
	}
	return points
}

func (p Point2) L1(p2 Point2) int {
	return AbsDiff(p.X, p2.X) + AbsDiff(p.Y, p2.Y)
}

func (p Point2) LInf(p2 Point2) int {
	return IntMax(AbsDiff(p.X, p2.X), AbsDiff(p.Y, p2.Y))
}

func (p Point2) L2(p2 Point2) float64 {
	return math.Sqrt(float64(AbsDiff(p.X, p2.X)*AbsDiff(p.X, p2.X) + AbsDiff(p.Y, p2.Y)*AbsDiff(p.Y, p2.Y)))
}

// Check if a point is inside a rectangle defined by two points
func PointInRectangle(p, p1, p2 Point2) bool {
	return (p.X >= IntMin(p1.X, p2.X) && p.X <= IntMax(p1.X, p2.X) &&
		p.Y >= IntMin(p1.Y, p2.Y) && p.Y <= IntMax(p1.Y, p2.Y))
}

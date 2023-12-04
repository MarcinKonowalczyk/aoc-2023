package utils

type Point struct {
	X int
	Y int
}

func NewPoint(x int, y int) Point {
	return Point{x, y}
}

func (p Point) Add(p2 Point) Point {
	return Point{p.X + p2.X, p.Y + p2.Y}
}

func (p Point) Sub(p2 Point) Point {
	return Point{p.X - p2.X, p.Y - p2.Y}
}

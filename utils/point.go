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

// Returns a slice of points in a 2D grid of the given width and height.
// Iterate over the points row by row, left to right, starting from 0,0.
func PointsIn2D(width, height int) []Point {
	points := make([]Point, 0, width*height)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			points = append(points, Point{x, y})
		}
	}
	return points
}

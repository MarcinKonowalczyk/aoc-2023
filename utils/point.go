package utils

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

package vector

type Point struct {
	X float64
	Y float64
}

func (p1 Point) Dot(p2 Point) float64 {
	return (p1.X * p2.X) + (p1.Y * p2.Y)
}

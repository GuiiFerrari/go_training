package algelin

import "math"

type Point2D struct {
	X   float64
	Y   float64
	Mod float64
}

func (p *Point2D) Modulus() {
	p.Mod = math.Sqrt(InnerProduct2d(p, p))
}

func NewPoint2D(x float64, y float64) *Point2D {
	p := &Point2D{X: x, Y: y}
	p.Modulus()
	return p
}

func InnerProduct2d(point_a *Point2D, point_b *Point2D) float64 {
	return point_a.X*point_b.X + point_a.Y*point_b.Y
}

type Vector2D struct {
	Origin *Point2D
	Versor *Point2D
}

func NewVector2D(origin *Point2D, versor *Point2D) *Vector2D {
	return &Vector2D{Origin: origin, Versor: versor}
}

func (l *Vector2D) PointAt(t float64) *Point2D {
	return NewPoint2D(
		l.Origin.X+t*l.Versor.X,
		l.Origin.Y+t*l.Versor.Y)
}

func AngleBetweenVectors2d(vector_a *Vector2D, vector_b *Vector2D) float64 {
	return InnerProduct2d(vector_a.Versor, vector_b.Versor) / (vector_a.Versor.Mod * vector_b.Versor.Mod)
}

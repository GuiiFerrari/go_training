package algelin

import "math"

type Point3D struct {
	X   float64
	Y   float64
	Z   float64
	Mod float64
}

func (p *Point3D) Modulus() {
	p.Mod = math.Sqrt(InnerProduct3d(*p, *p))
}

func NewPoint3D(x float64, y float64, z float64) *Point3D {
	p := &Point3D{X: x, Y: y, Z: z}
	p.Modulus()
	return p
}

func InnerProduct3d(point_a Point3D, point_b Point3D) float64 {
	return point_a.X*point_b.X + point_a.Y*point_b.Y + point_a.Z*point_b.Z
}

type Vector3D struct {
	Origin *Point3D
	Versor *Point3D
}

func CrossProduct3d(vector_a *Vector3D, vector_b *Vector3D) *Point3D {
	return NewPoint3D(
		vector_a.Versor.Y*vector_b.Versor.Z-vector_a.Versor.Z*vector_b.Versor.Y,
		vector_a.Versor.Z*vector_b.Versor.X-vector_a.Versor.X*vector_b.Versor.Z,
		vector_a.Versor.X*vector_b.Versor.Y-vector_a.Versor.Y*vector_b.Versor.X)
}

func NormalizedVector3d(point_a *Point3D) *Point3D {
	mod := point_a.Mod
	return NewPoint3D(point_a.X/mod, point_a.Y/mod, point_a.Z/mod)
}

func AngleBetweenVectors3d(point_a *Point3D, point_b *Point3D) float64 {
	return InnerProduct3d(*point_a, *point_b) / (point_a.Mod * point_b.Mod)
}

func NewVector3D(origin *Point3D, direction *Point3D) *Vector3D {
	return &Vector3D{Origin: origin, Versor: direction}
}

func (l *Vector3D) PointAt(t float64) *Point3D {
	return NewPoint3D(
		l.Origin.X+t*l.Versor.X,
		l.Origin.Y+t*l.Versor.Y,
		l.Origin.Z+t*l.Versor.Z)
}

// func

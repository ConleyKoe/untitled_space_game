package main

import "math"

type Vec3 struct {
	X, Y, Z float64
}

func (a *Vec3) Add(b Vec3) Vec3 { //Adds vector b to vector a, called like this: c := a.Add(b)
	return Vec3{
		X: a.X + b.X,
		Y: a.Y + b.Y,
		Z: a.Z + b.Z}
}

func (a *Vec3) Subtract(b Vec3) Vec3 { //Subtracts vector a by vector b, called like: c := a.Subtract(b)
	return Vec3{
		X: a.X - b.X,
		Y: a.Y - b.Y,
		Z: a.Z - b.Z,
	}
}

func (a *Vec3) Scale(b float64) Vec3 { //Multiplies vector a by float b, called like: c := a.Scale(b)
	return Vec3{
		X: a.X * b,
		Y: a.Y * b,
		Z: a.Z * b,
	}
}

func (a *Vec3) Normalize() Vec3 { //Normalizes vector a, called like: b := a.Normalize()
	m := math.Sqrt(a.X*a.X + a.Y*a.Y + a.Z*a.Z) //Calculates the magnitude of the vector
	if m == 0 {
		return Vec3{0, 0, 0} //If the magnitude is zero, return the zero vector
	}
	return Vec3{
		X: a.X / m,
		Y: a.Y / m,
		Z: a.Z / m,
	}
}

func (a *Vec3) DotProduct(b Vec3) float64 { //Finds the dot product between the a and b vectors, called like: c := a.DotProduct(b)
	return (a.X * b.X) + (a.Y * b.Y) + (a.Z * b.Z)
}

func (a *Vec3) CrossProduct(b Vec3) Vec3 { //Finds the cross product between the a and b vectors, called like: c := a.CrossProduct(b)
	return Vec3{
		X: (a.Y * b.Z) - (a.Z * b.Y),
		Y: (a.Z * b.X) - (a.X * b.Z),
		Z: (a.X * b.Y) - (a.Y * b.X),
	}
}

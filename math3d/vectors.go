package math3d

import "math"

type Vec3 struct {
	X, Y, Z float64
}

type Vec2 struct {
	X, Y float64
}

type Vec2Int struct {
	X, Y int
}

type Vec4 struct {
	X, Y, Z, W float64
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

func (a Vec3) Scale(b float64) Vec3 { //Multiplies vector a by float b, called like: c := a.Scale(b)
	return Vec3{
		X: a.X * b,
		Y: a.Y * b,
		Z: a.Z * b,
	}
}

func (a Vec3) Normalize() Vec3 { //Normalizes vector a, called like: b := a.Normalize()
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

func (a Vec3) ToVec4(w float64) Vec4 {
	return Vec4{a.X, a.Y, a.Z, w}
}

func (v Vec4) ToVec3() Vec3 {
	return Vec3{v.X, v.Y, v.Z}
}

// Add returns the component-wise addition of v and u.
func (v Vec4) Add(u Vec4) Vec4 {
	return Vec4{v.X + u.X, v.Y + u.Y, v.Z + u.Z, v.W + u.W}
}

// Sub returns the component-wise subtraction of u from v.
func (v Vec4) Sub(u Vec4) Vec4 {
	return Vec4{v.X - u.X, v.Y - u.Y, v.Z - u.Z, v.W - u.W}
}

// Scale returns the vector v multiplied by scalar s.
func (v Vec4) Scale(s float64) Vec4 {
	return Vec4{v.X * s, v.Y * s, v.Z * s, v.W * s}
}

// Dot returns the dot product of v and u.
// Usually, for geometric purposes, we exclude W from the dot product.
func (v Vec4) Dot(u Vec4) float64 {
	return v.X*u.X + v.Y*u.Y + v.Z*u.Z
}

// Length returns the length (magnitude) of the vector, ignoring W.
func (v Vec4) Length() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

// Normalize returns a normalized version of v (only X,Y,Z normalized, W unchanged).
func (v Vec4) Normalize() Vec4 {
	length := v.Length()
	if length == 0 {
		return v
	}
	return Vec4{v.X / length, v.Y / length, v.Z / length, v.W}
}

// PerspectiveDivide divides X, Y, Z by W.
// Returns zero vector if W is zero to avoid division by zero.
func (v Vec4) PerspectiveDivide() Vec4 {
	if v.W == 0 {
		return Vec4{0, 0, 0, 0}
	}
	invW := 1 / v.W
	return Vec4{v.X * invW, v.Y * invW, v.Z * invW, 1}
}

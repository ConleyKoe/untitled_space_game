package math3d

import (
	"math"
)

type Quaternion struct { //some sort of 4d vector which stores an axis (x, y, z) and an angle of rotation (w)
	W, X, Y, Z float64
}

func (q Quaternion) Multiply(b Quaternion) Quaternion { //to be completely honest, i have no idea how these things work, i'm just following a tutorial...
	return Quaternion{
		W: q.W*b.W - q.X*b.X - q.Y*b.Y - q.Z*b.Z,
		X: q.W*b.X + q.X*b.W + q.Y*b.Z - q.Z*b.Y,
		Y: q.W*b.Y - q.X*b.Z + q.Y*b.W + q.Z*b.X,
		Z: q.W*b.Z + q.X*b.Y - q.Y*b.X + q.Z*b.W,
	}
}

func (q Quaternion) Normalize() Quaternion { //essentially the same math as math3d.Vec3.Normalize(), just with an extra term
	m := math.Sqrt(q.W*q.W + q.X*q.X + q.Y*q.Y + q.Z*q.Z) //calculates the magnitude of the quaternion
	return Quaternion{                                    //returns a quaternion with each term normalized
		W: q.W / m,
		X: q.X / m,
		Y: q.Y / m,
		Z: q.Z / m,
	}
}

func NewQuatFromAxisAngle(axis Vec3, angle float64) Quaternion { //creates a new quaternion from a given axis and angle
	halfAngle := angle / 2         //something somethin...need a half angle for the formula
	sinHalf := math.Sin(halfAngle) //also need the sin of angle/2
	return Quaternion{             //uses the quaternion formula to make a new one from our axis and the half angle
		W: math.Cos(halfAngle),
		X: axis.X * sinHalf,
		Y: axis.Y * sinHalf,
		Z: axis.Z * sinHalf,
	}.Normalize() //normalizes said quaternion
}

func (q Quaternion) RotateVector(v Vec3) Vec3 { //rotates vector v according to quaternion q
	qVec := Quaternion{0, v.X, v.Y, v.Z}       //creates a new quaternion with axis of v and angle of 0
	qConj := Quaternion{q.W, -q.X, -q.Y, -q.Z} //essentially a copy of q

	result := q.Multiply(qVec).Multiply(qConj) //multiplies our vector quaternion with q
	return Vec3{result.X, result.Y, result.Z}  //returns the newly rotated vector
}

func (q Quaternion) ToMatrix() Mat4 { //creates a new matrix from our quaternion q, looks a mess but should work...don't ask me how
	xx := q.X * q.X
	yy := q.Y * q.Y
	zz := q.Z * q.Z
	xy := q.X * q.Y
	xz := q.X * q.Z
	yz := q.Y * q.Z
	wx := q.W * q.X
	wy := q.W * q.Y
	wz := q.W * q.Z

	return Mat4{
		{1 - 2*(yy+zz), 2 * (xy - wz), 2 * (xz + wy), 0},
		{2 * (xy + wz), 1 - 2*(xx+zz), 2 * (yz - wx), 0},
		{2 * (xz - wy), 2 * (yz + wx), 1 - 2*(xx+yy), 0},
		{0, 0, 0, 1},
	}
}

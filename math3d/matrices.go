package math3d

import "math"

type Mat4 [4][4]float64

func (m Mat4) Transpose() Mat4 { //gives the inverse of matrix m
	var result Mat4
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			result[i][j] = m[j][i]
		}
	}
	return result
}

func Translate(v Vec3) Mat4 { //creates a translation matrix to translate vector v
	return Mat4{
		{1, 0, 0, v.X},
		{0, 1, 0, v.Y},
		{0, 0, 1, v.Z},
		{0, 0, 0, 1},
	}
}

func (a Mat4) Multiply(b Mat4) Mat4 { //multiplies matrix a by matrix b
	var result Mat4

	for row := 0; row < 4; row++ {
		for col := 0; col < 4; col++ {
			sum := 0.0
			for k := 0; k < 4; k++ {
				sum += a[row][k] * b[k][col]
			}
			result[row][col] = sum
		}
	}

	return result
}

func BuildModelMatrix(position, scale Vec3, rotation Quaternion) Mat4 { //Builds a model matrix to transform an object from local space into world space
	return Translate(position).Multiply(rotation.ToMatrix()).Multiply(Scale(scale))
}

func Scale(s Vec3) Mat4 {
	return Mat4{
		{s.X, 0, 0, 0},
		{0, s.Y, 0, 0},
		{0, 0, s.Z, 0},
		{0, 0, 0, 1},
	}
}

func BuildProjectionMatrix(fov, aspect, near, far float64) Mat4 {
	f := 1.0 / math.Tan(fov/2)

	return Mat4{
		{f / aspect, 0, 0, 0},
		{0, f, 0, 0},
		{0, 0, (far + near) / (near - far), (2 * far * near) / (near - far)},
		{0, 0, -1, 0},
	}
}

func (m Mat4) MulVec4(v Vec4) Vec4 { //multiplies matrix m by vector4 v
	return Vec4{
		X: m[0][0]*v.X + m[0][1]*v.Y + m[0][2]*v.Z + m[0][3]*v.W,
		Y: m[1][0]*v.X + m[1][1]*v.Y + m[1][2]*v.Z + m[1][3]*v.W,
		Z: m[2][0]*v.X + m[2][1]*v.Y + m[2][2]*v.Z + m[2][3]*v.W,
		W: m[3][0]*v.X + m[3][1]*v.Y + m[3][2]*v.Z + m[3][3]*v.W,
	}
}

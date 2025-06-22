package camera

import (
	"untitled_space_game/math3d"
)

type ViewMatrix [4][4]float64

type Camera struct {
	Position math3d.Vec3
	Forward  math3d.Vec3 //{0,0,-1}
	Up       math3d.Vec3 //{0,1,0}
	Right    math3d.Vec3 //{1,0,0}
}

func BuildViewMatrix(camera Camera) ViewMatrix {
	f := camera.Forward.Normalize()

	// Ensure right and up are orthogonal
	r := f.CrossProduct(camera.Up)
	r = r.Normalize()
	u := r.CrossProduct(f)
	u = u.Normalize() //still haven't fixed this but i don't know what's causing the problem

	p := camera.Position

	return ViewMatrix{
		{r.X, r.Y, r.Z, -r.DotProduct(p)},
		{u.X, u.Y, u.Z, -u.DotProduct(p)},
		{-f.X, -f.Y, -f.Z, f.DotProduct(p)}, // Forward is inverted
		{0, 0, 0, 1},
	}
}

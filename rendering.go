package main

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

func drawLine(screen *ebiten.Image, x1, y1, x2, y2 float64, clr color.Color) {
	x0 := int(x1)   // converts x1 to an integer so it can be used in our loop
	y0 := int(y1)   // this ^ but for y1
	endX := int(x2) //literally the same...
	endY := int(y2) //once again...

	deltaX := int(math.Abs(x2 - x1))
	deltaY := int(-math.Abs(y2 - y1))

	if endX == endY || endY == endX {
		return
	}

	stepX := 1 //direction the algo steps along the x axis
	if x0 > endX {
		stepX = -1
	} //if slope is negative, move left along x instead of right, otherwise stepX remains at 1 (right)
	stepY := 1 //direction the algo steps along the y axis
	if y0 > endY {
		stepY = -1
	} //if the line goes upwards, move up along y instead of down, otherwise stepY remains at 1 (down)

	err := deltaX + deltaY //initializes the error term as the delta of X plus the delta of Y, to be multiplied by 2 later

	for {

		if x0 >= 0 && x0 < screen.Bounds().Dx() && y0 >= 0 && y0 < screen.Bounds().Dy() {
			screen.Set(x0, y0, clr)
		}

		if x0 == endX && y0 == endY { //stops drawing the line once we've reached the endpoint
			break
		}
		err2 := 2 * err //properly sets up the error term for drawing

		if err2 >= deltaY { //moves along x by stepX if needed, then changes the error term
			err += deltaY
			x0 += stepX
		}

		if err2 <= deltaX { //moves along y by stepY if needed, then changes the error term
			err += deltaX
			y0 += stepY
		}

	}
}

func ProjectPoint(v Vec3, focalLength, screenWidth, screenHeight float64) Vec2 { //Projects 3d vector v to a 2d vector
	scale := focalLength / v.Z
	x := v.X*scale + screenWidth/2
	y := v.Y*scale + screenHeight/2

	return Vec2{x, y}
}

func RenderEdge(screen *ebiten.Image, a, b Vec3, clr color.Color) {
	if a.Z <= 0.1 || b.Z <= 0.1 { // Avoid near-zero Z or points behind camera
		return
	}

	p1 := ProjectPoint(a, 100, 640, 480)
	p2 := ProjectPoint(b, 100, 640, 480)

	// Sanity check: Don't draw if either point is NaN or infinite
	if math.IsNaN(float64(p1.X)) || math.IsNaN(float64(p1.Y)) || math.IsNaN(float64(p2.X)) || math.IsNaN(float64(p2.Y)) {
		return
	}
	if math.IsInf(p1.X, 0) || math.IsInf(p1.Y, 0) || math.IsInf(p2.X, 0) || math.IsInf(p2.Y, 0) {
		return
	}
	drawLine(screen, math.Round(p1.X), math.Round(p1.Y), math.Round(p2.X), math.Round(p2.Y), clr)
}

package main

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

func drawLine(screen *ebiten.Image, x1, y1, x2, y2 float64, clr color.Color) {
	deltaX := math.Abs(x2 - x1)
	deltaY := math.Abs(y2 - y1)

	stepX := 1 //direction the algo steps along the x axis
	if x1 > x2 {
		stepX = -1
	} //if slope is negative, move left along x instead of right, otherwise stepX remains at 1 (right)
	stepY := 1 //direction the algo steps along the y axis
	if y1 > y2 {
		stepY = -1
	} //if the line goes upwards, move up along y instead of down, otherwise stepY remains at 1 (down)

	err := deltaX - deltaY //initializes the error term as the delta of X minus the delta of Y, to be multiplied by 2 later

	x0 := int(x1)   // converts x1 to an integer so it can be used in our loop
	y0 := int(y1)   // this ^ but for y1
	endX := int(x2) //literally the same...
	endY := int(y2) //once again...
	for {

		screen.Set(x0, y0, clr)       //plots the first point of the line in the specified color
		if x0 == endX && y0 == endY { //stops drawing the line once we've reached the endpoint
			break
		}
		err2 := 2 * err //properly sets up the error term for drawing

		if err2 > -deltaY { //moves along x by stepX if needed, then changes the error term
			err = err - deltaY
			x0 = x0 + stepX
		}

		if err2 < deltaX { //moves along y by stepY if needed, then changes the error term
			err = err + deltaX
			y0 = y0 + stepY
		}

	}
}

func ProjectPoint(v Vec3, focalLength, screenWidth, screenHeight float64) Vec2 { //Projects 3d vector v to a 2d vector
	scale := focalLength / v.Z
	x := v.X*scale + screenWidth/2
	y := v.Y*scale + screenHeight/2

	return Vec2{x, y}
}

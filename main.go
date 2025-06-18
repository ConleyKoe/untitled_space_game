package main

import (
	"fmt"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

var Vertices = []Vec3{{1, 1, 1}, {2, 1, 1}, {1, 2, 1}}

type Game struct {
	focalLength float64
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	p1 := ProjectPoint(Vertices[0], g.focalLength, 640, 480)
	fmt.Println(p1)
	p2 := ProjectPoint(Vertices[1], g.focalLength, 640, 480)
	fmt.Println(p2)
	p3 := ProjectPoint(Vertices[2], g.focalLength, 640, 480)
	fmt.Println(p3)

	drawLine(screen, p1.X, p1.Y, p2.X, p2.Y, color.White)
	drawLine(screen, p1.X, p1.Y, p3.X, p3.Y, color.White)
	drawLine(screen, p2.X, p2.Y, p3.X, p3.Y, color.White)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(&Game{
		focalLength: 100,
	}); err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

var TestMesh = Mesh{
	Vertices: []Vec3{
		{1, 1, 1}, {2, 1, 1}, {1, 2, 1}},
	Faces: []Face{
		{0, 1, 2},
	},
	Edges: []Edge{
		{0, 1}, {0, 2},
	},
}

var TestObject = MeshObject{
	Mesh:     &TestMesh,
	Position: Vec3{1, 1, 1},
}

type Game struct {
	focalLength float64
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.Black)
	TestObject.DrawMeshObject(screen, color.White)

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("")
	if err := ebiten.RunGame(&Game{
		focalLength: 100,
	}); err != nil {
		log.Fatal(err)
	}
}

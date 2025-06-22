package main

import (
	"image/color"
	"log"

	"untitled_space_game/math3d"
	"untitled_space_game/rendering"

	"github.com/hajimehoshi/ebiten/v2"
)

var TestMesh = rendering.Mesh{
	Vertices: []math3d.Vec3{
		{X: -2, Y: -2, Z: -2}, // 0
		{X: 2, Y: -2, Z: -2},  // 1
		{X: 2, Y: 2, Z: -2},   // 2
		{X: -2, Y: 2, Z: -2},  // 3
		{X: -2, Y: -2, Z: 2},  // 4
		{X: 2, Y: -2, Z: 2},   // 5
		{X: 2, Y: 2, Z: 2},    // 6
		{X: -2, Y: 2, Z: 2},   // 7
	},
	Faces: []rendering.Face{
		// Front face (+Z)
		{A: 4, B: 5, C: 6}, {A: 4, B: 6, C: 7},

		// Back face (-Z)
		{A: 0, B: 2, C: 1}, {A: 0, B: 3, C: 2},

		// Right face (+X)
		{A: 1, B: 2, C: 6}, {A: 1, B: 6, C: 5},

		// Left face (-X)
		{A: 0, B: 7, C: 3}, {A: 0, B: 4, C: 7},

		// Top face (+Y)
		{A: 3, B: 7, C: 6}, {A: 3, B: 6, C: 2},

		// Bottom face (-Y)
		{A: 0, B: 1, C: 5}, {A: 0, B: 5, C: 4},
	},
	Edges: []rendering.Edge{
		//Front of mesh
		{A: 4, B: 5, AdjacentFaces: [2]int{0, 11}}, {A: 5, B: 6, AdjacentFaces: [2]int{0, 5}},
		{A: 4, B: 7, AdjacentFaces: [2]int{1, 7}}, {A: 6, B: 7, AdjacentFaces: [2]int{1, 8}},
		//Back of mesh
		{A: 2, B: 1, AdjacentFaces: [2]int{2, 4}}, {A: 0, B: 1, AdjacentFaces: [2]int{2, 10}},
		{A: 0, B: 3, AdjacentFaces: [2]int{3, 6}}, {A: 3, B: 2, AdjacentFaces: [2]int{3, 9}},
		//Right side of mesh
		{A: 2, B: 6, AdjacentFaces: [2]int{4, 9}}, {A: 1, B: 5, AdjacentFaces: [2]int{5, 10}},
		//Left side of mesh
		{A: 3, B: 7, AdjacentFaces: [2]int{8, 9}}, {A: 0, B: 4, AdjacentFaces: [2]int{7, 11}},
	},
}

var TestObject = rendering.MeshObject{
	Mesh:     &TestMesh,
	Position: math3d.Vec3{X: 0, Y: -6, Z: 10},
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

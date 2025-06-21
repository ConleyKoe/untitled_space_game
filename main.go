package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

var TestMesh = Mesh{
	Vertices: []Vec3{
		{-2, -2, -2}, // 0
		{2, -2, -2},  // 1
		{2, 2, -2},   // 2
		{-2, 2, -2},  // 3
		{-2, -2, 2},  // 4
		{2, -2, 2},   // 5
		{2, 2, 2},    // 6
		{-2, 2, 2},   // 7
	},
	Faces: []Face{
		// Front face (+Z)
		{4, 5, 6}, {4, 6, 7},

		// Back face (-Z)
		{0, 2, 1}, {0, 3, 2},

		// Right face (+X)
		{1, 2, 6}, {1, 6, 5},

		// Left face (-X)
		{0, 7, 3}, {0, 4, 7},

		// Top face (+Y)
		{3, 7, 6}, {3, 6, 2},

		// Bottom face (-Y)
		{0, 1, 5}, {0, 5, 4},
	},
	Edges: []Edge{
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

var TestObject = MeshObject{
	Mesh:     &TestMesh,
	Position: Vec3{0, 0, 4},
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

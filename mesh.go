package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Edge struct {
	A, B int
}

type Mesh struct {
	Vertices []Vec3
	Faces    []Vec3
	Edges    []Edge
}

type MeshObject struct {
	Mesh     Mesh
	Position Vec3
}

func (Object *MeshObject) DrawMeshObject(screen *ebiten.Image, clr color.Color) {
	for _, i := range Object.Mesh.Edges {
		a := Object.Mesh.Vertices[i.A].Add(Object.Position)
		b := Object.Mesh.Vertices[i.B].Add(Object.Position)
		RenderEdge(screen, a, b, clr)
	}
}

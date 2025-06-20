package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Edge struct { //An edge holds two integers which act as indices to a mesh's vertices, denotes the edges of a mesh
	A, B int
}

type Face struct { //A face holds three integers which act as indices to a mesh's vertices, denotes the polygons of a mesh
	A, B, C int
}

type Mesh struct { //Holds a list of 3d vertices, a list of polygons, and a list of edges to be drawn around said polygons
	Vertices []Vec3
	Faces    []Face //These will just be used for backface culling so the renderer will know which edges to draw
	Edges    []Edge //Only edges are actually drawn
}

type MeshObject struct { //The physical manifestation of a mesh, holds a pointer to a mesh and a position vector in world space
	Mesh     *Mesh
	Position Vec3
}

func (Object *MeshObject) DrawMeshObject(screen *ebiten.Image, clr color.Color) { //Draws all the edges of a mesh object, called like: Object.DrawMeshObject(screen, color.White)
	for _, i := range Object.Mesh.Edges { //Interates through the mesh's edges
		a := Object.Mesh.Vertices[i.A].Add(Object.Position) //Adds the mesh object's position vector to the vertex specified in the mesh's edge
		b := Object.Mesh.Vertices[i.B].Add(Object.Position)
		RenderEdge(screen, a, b, clr) //This function will project the two vectors into 2d and draw a line between them
	}
}

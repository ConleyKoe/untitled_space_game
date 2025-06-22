package rendering

import (
	"image/color"
	"untitled_space_game/math3d"

	"github.com/hajimehoshi/ebiten/v2"
)

type Edge struct { //An edge holds two integers which act as indices to a mesh's vertices, denotes the edges of a mesh
	A, B          int
	AdjacentFaces [2]int
}

type Face struct { //A face holds three integers which act as indices to a mesh's vertices, denotes the polygons of a mesh
	A, B, C int
}

type Mesh struct { //Holds a list of 3d vertices, a list of polygons, and a list of edges to be drawn around said polygons
	Vertices []math3d.Vec3
	Faces    []Face //These will just be used for backface culling so the renderer will know which edges to draw
	Edges    []Edge //Only edges are actually drawn
}

type MeshObject struct { //The physical manifestation of a mesh, holds a pointer to a mesh and a position vector in world space
	Mesh     *Mesh
	Position math3d.Vec3
}

func (Mesh *Mesh) IsFaceVisible(a Face) bool { //Calculates the surface normal of a given edge and returns visibility info as a boolean
	normal := (Mesh.Vertices[a.B].Subtract(Mesh.Vertices[a.A]))
	newNormal := normal.CrossProduct(Mesh.Vertices[a.C].Subtract(Mesh.Vertices[a.A]))
	normNormal := newNormal.Normalize() //Using three different variables because for some reason I can't call crossproduct and normalize in the same variable? idk i'll fix it later probably
	return normNormal.DotProduct(Mesh.Vertices[a.A]) > 0
} //MeshObject.Mesh.IsFaceVisible(MeshObject.Mesh.Faces[0])

func (Object *MeshObject) DrawMeshObject(screen *ebiten.Image, clr color.Color) { //Draws all the edges of a mesh object, called like: Object.DrawMeshObject(screen, color.White)
	for _, i := range Object.Mesh.Edges { //Interates through the mesh's edges
		//There's probably a simpler way to do this, but i don't feel like figuring it out
		if Object.Mesh.IsFaceVisible(Object.Mesh.Faces[i.AdjacentFaces[0]]) || Object.Mesh.IsFaceVisible(Object.Mesh.Faces[i.AdjacentFaces[1]]) { //If either of the adjacent faces of the current edge is visible, transform the endpoints and render the edge
			a := Object.Mesh.Vertices[i.A].Add(Object.Position) //Adds the mesh object's position vector to the vertex specified in the mesh's edge
			b := Object.Mesh.Vertices[i.B].Add(Object.Position)
			RenderEdge(screen, a, b, clr) //This function will project the two vectors into 2d and draw a line between them
		} else {
			continue
		}

	}
}

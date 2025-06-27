package rendering

import (
	"image/color"
	"untitled_space_game/camera"
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
	Rotation math3d.Quaternion
}

func (Mesh *Mesh) IsFaceVisible(a [3]math3d.Vec3, cameraPos math3d.Vec3) bool {
	// Calculate the surface normal of the face
	edge1 := a[1].Subtract(a[0])
	edge2 := a[2].Subtract(a[0])
	normal := edge1.CrossProduct(edge2).Normalize()

	// Calculate the view vector (from camera to the face)
	viewVector := a[0].Subtract(cameraPos).Normalize()

	// Backface culling: if dot product is less than zero, face is visible
	return normal.DotProduct(viewVector) < 0
}

func (Object *MeshObject) DrawMeshObject(screen *ebiten.Image, cam camera.Camera, clr color.Color) { //Draws all the edges of a mesh object, called like: Object.DrawMeshObject(screen, color.White)
	modelMatrix := math3d.BuildModelMatrix(Object.Position, Object.Rotation, math3d.Vec3{X: 1, Y: 1, Z: 1}) //creates the model matrix based on the attributes of the meshobject
	viewMatrix := camera.BuildViewMatrix(cam)                                                               //creates the camera view matrix
	projectionMatrix := math3d.BuildProjectionMatrix(80, 640/480, 0.01, 4000)                               //creates the projection matrix
	MVP := projectionMatrix.Multiply(viewMatrix).Multiply(modelMatrix)                                      //combines all three

	for _, i := range Object.Mesh.Edges { //Interates through the mesh's edges
		//There's probably a simpler way to do this, but i don't feel like figuring it out
		face1 := Object.Mesh.Faces[i.AdjacentFaces[0]]
		face2 := Object.Mesh.Faces[i.AdjacentFaces[1]]
		mmFace1 := [3]math3d.Vec3{
			modelMatrix.MulVec4(Object.Mesh.Vertices[face1.A].ToVec4(1)).ToVec3(),
			modelMatrix.MulVec4(Object.Mesh.Vertices[face1.B].ToVec4(1)).ToVec3(),
			modelMatrix.MulVec4(Object.Mesh.Vertices[face1.C].ToVec4(1)).ToVec3(),
		}
		mmFace2 := [3]math3d.Vec3{
			modelMatrix.MulVec4(Object.Mesh.Vertices[face2.A].ToVec4(1)).ToVec3(),
			modelMatrix.MulVec4(Object.Mesh.Vertices[face2.B].ToVec4(1)).ToVec3(),
			modelMatrix.MulVec4(Object.Mesh.Vertices[face2.C].ToVec4(1)).ToVec3(),
		}

		if Object.Mesh.IsFaceVisible(mmFace1, cam.Position) || Object.Mesh.IsFaceVisible(mmFace2, cam.Position) { //If either of the adjacent faces of the current edge is visible, transform the endpoints and render the edge
			a := math3d.Vec4{X: Object.Mesh.Vertices[i.A].X, Y: Object.Mesh.Vertices[i.A].Y, Z: Object.Mesh.Vertices[i.A].Z, W: 1.0} //converts our 3d points to 4d vectors to be used in our matrices
			b := math3d.Vec4{X: Object.Mesh.Vertices[i.B].X, Y: Object.Mesh.Vertices[i.B].Y, Z: Object.Mesh.Vertices[i.B].Z, W: 1.0}
			a2 := MVP.MulVec4(a)
			b2 := MVP.MulVec4(b)
			aF := math3d.Vec3{X: a2.X / a2.W, Y: a2.Y / a2.W, Z: a2.Z / a2.W}
			bF := math3d.Vec3{X: b2.X / b2.W, Y: b2.Y / b2.W, Z: b2.Z / b2.W}
			RenderEdge(screen, aF, bF, clr) //This function will project the two vectors into 2d and draw a line between them
		} else {
			continue
		}

	}
}

func (Object *MeshObject) DrawMeshObjectFaces(screen *ebiten.Image, cam camera.Camera, clr color.Color) {
	modelMatrix := math3d.BuildModelMatrix(Object.Position, Object.Rotation, math3d.Vec3{X: 1, Y: 1, Z: 1}) //creates the model matrix based on the attributes of the meshobject
	viewMatrix := camera.BuildViewMatrix(cam)                                                               //creates the camera view matrix
	projectionMatrix := math3d.BuildProjectionMatrix(80, 640/480, 0.01, 4000)                               //creates the projection matrix
	MVP := projectionMatrix.Multiply(viewMatrix).Multiply(modelMatrix)                                      //combines all three

	for _, i := range Object.Mesh.Faces {
		transTri := [3]math3d.Vec3{
			modelMatrix.MulVec4(Object.Mesh.Vertices[i.A].ToVec4(1)).ToVec3(),
			modelMatrix.MulVec4(Object.Mesh.Vertices[i.B].ToVec4(1)).ToVec3(),
			modelMatrix.MulVec4(Object.Mesh.Vertices[i.C].ToVec4(1)).ToVec3(),
		}
		if Object.Mesh.IsFaceVisible(transTri, cam.Position) {
			drawTri := [3]math3d.Vec3{
				MVP.MulVec4(Object.Mesh.Vertices[i.A].ToVec4(1)).PerspectiveDivide().ToVec3(),
				MVP.MulVec4(Object.Mesh.Vertices[i.B].ToVec4(1)).PerspectiveDivide().ToVec3(),
				MVP.MulVec4(Object.Mesh.Vertices[i.C].ToVec4(1)).PerspectiveDivide().ToVec3(),
			}
			//fmt.Println("drawing face", i)
			RenderFace(screen, drawTri[0], drawTri[1], drawTri[2], clr)
		}
	}

}

package camera

import (
	"untitled_space_game/math3d"

	"github.com/hajimehoshi/ebiten/v2"
)

type Camera struct {
	Position math3d.Vec3
	Rotation math3d.Quaternion
}

func BuildViewMatrix(camera Camera) math3d.Mat4 { //builds a view matrix from the camera's rotation quaternion using some complicated formula
	return camera.Rotation.ToMatrix().Transpose().Multiply(math3d.Translate(camera.Position.Scale(-1)))
}

// these functions calculate the camera's direction vectors from the rotation quaternion
func (cam *Camera) Forward() math3d.Vec3 {
	return cam.Rotation.RotateVector(math3d.Vec3{X: 0, Y: 0, Z: 1}).Normalize()
}

func (cam *Camera) Right() math3d.Vec3 {
	return cam.Rotation.RotateVector(math3d.Vec3{X: 1, Y: 0, Z: 0}).Normalize()
}

func (cam *Camera) Up() math3d.Vec3 {
	return cam.Rotation.RotateVector(math3d.Vec3{X: 0, Y: 1, Z: 0}).Normalize()
}

func (cam *Camera) handleMovement() {
	moveSpeed := 0.1 // example speed

	if ebiten.IsKeyPressed(ebiten.KeyW) {
		cam.Position = cam.Position.Add(cam.Forward().Scale(moveSpeed))
	}

	if ebiten.IsKeyPressed(ebiten.KeyS) {
		cam.Position = cam.Position.Add(cam.Forward().Scale(-moveSpeed))
	}

	if ebiten.IsKeyPressed(ebiten.KeyA) {
		cam.Position = cam.Position.Add(cam.Right().Scale(moveSpeed))
	}

	if ebiten.IsKeyPressed(ebiten.KeyD) {
		cam.Position = cam.Position.Add(cam.Right().Scale(-moveSpeed))
	}
}

func (camera *Camera) handleRotation() {
	rotationSpeed := 0.02

	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		delta := math3d.NewQuatFromAxisAngle(camera.Up(), rotationSpeed) //finds the quaternion to apply a rotation around the up vector
		camera.Rotation = delta.Multiply(camera.Rotation).Normalize()
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		delta := math3d.NewQuatFromAxisAngle(camera.Up(), -rotationSpeed) //opposite
		camera.Rotation = delta.Multiply(camera.Rotation).Normalize()
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		delta := math3d.NewQuatFromAxisAngle(camera.Right(), rotationSpeed) //finds the quaternion to apply a rotations around the right vector
		camera.Rotation = delta.Multiply(camera.Rotation).Normalize()
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		delta := math3d.NewQuatFromAxisAngle(camera.Right(), -rotationSpeed) //opposite
		camera.Rotation = delta.Multiply(camera.Rotation).Normalize()
	}

}

func (camera *Camera) Update() {
	camera.handleMovement()
	camera.handleRotation()
}

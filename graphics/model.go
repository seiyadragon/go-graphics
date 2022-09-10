package graphics

import (
	"github.com/go-gl/mathgl/mgl32"
)

type Model struct {
	Mesh     Mesh
	Material Material
	Position mgl32.Vec3
	Rotation mgl32.Vec3
	Scale    mgl32.Vec3
}

func NewModel(mesh Mesh, material Material, position mgl32.Vec3, rotation mgl32.Vec3, scale mgl32.Vec3) Model {
	return Model{mesh, material, position, rotation, scale}
}

func (m Model) GetModel() mgl32.Mat4 {
	transform := mgl32.Translate3D(m.Position.X(), m.Position.Y(), m.Position.Z())
	scale := mgl32.Scale3D(m.Scale.X(), m.Scale.Y(), m.Scale.Z())
	rotate := mgl32.Rotate3DX(m.Rotation.X()).Mul3(mgl32.Rotate3DY(m.Rotation.Y()).Mul3(mgl32.Rotate3DZ(m.Rotation.Z())))

	return transform.Mul4(scale.Mul4(rotate.Mat4()))
}

func (m Model) Draw() {
	m.Mesh.Draw(m.Material)
}

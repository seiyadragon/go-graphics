package graphics

import "github.com/go-gl/mathgl/mgl32"

type View struct {
	Width    float32
	Height   float32
	Position mgl32.Vec3
	Front    mgl32.Vec3
	Up       mgl32.Vec3
}

func NewView(width float32, height float32, position mgl32.Vec3, target mgl32.Vec3, up mgl32.Vec3) View {
	return View{width, height, position, target, up}
}

func (v View) GetView() mgl32.Mat4 {
	return mgl32.Perspective(mgl32.DegToRad(45.0), v.Width/v.Height, 0.1, 100.0).Mul4(mgl32.LookAtV(v.Position, v.Position.Add(v.Front), v.Up))
}

type Layer struct {
	View     View
	Position mgl32.Vec3
	Rotation mgl32.Vec3
	Scale    mgl32.Vec3
}

func NewLayer(view View, position mgl32.Vec3, rotation mgl32.Vec3, scale mgl32.Vec3) Layer {
	return Layer{view, position, rotation, scale}
}

func DefaultLayer(view View) Layer {
	return NewLayer(view, mgl32.Vec3{0, 0, 0}, mgl32.Vec3{0, 0, 0}, mgl32.Vec3{1, 1, 1})
}

func (m Layer) GetLayer() mgl32.Mat4 {
	transform := mgl32.Translate3D(m.Position.X(), m.Position.Y(), m.Position.Z())
	scale := mgl32.Scale3D(m.Scale.X(), m.Scale.Y(), m.Scale.Z())
	rotate := mgl32.Rotate3DX(m.Rotation.X()).Mul3(mgl32.Rotate3DY(m.Rotation.Y()).Mul3(mgl32.Rotate3DZ(m.Rotation.Z())))

	return transform.Mul4(scale.Mul4(rotate.Mat4()))
}

func (m Layer) Draw(drawable Drawable) {
	drawable.Draw(m)
}

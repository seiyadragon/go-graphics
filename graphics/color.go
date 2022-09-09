package graphics

import "github.com/go-gl/mathgl/mgl32"

type Color struct {
	R uint8
	G uint8
	B uint8
	A uint8
}

func NewColor(r uint8, g uint8, b uint8, a uint8) Color {
	return Color{r, g, b, a}
}

func (color Color) ToVec4() mgl32.Vec4 {
	return mgl32.Vec4{float32(color.R) / 255, float32(color.G) / 255, float32(color.B) / 255, float32(color.A) / 255}
}

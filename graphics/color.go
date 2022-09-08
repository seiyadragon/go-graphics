package graphics

type Color struct {
	R uint8
	G uint8
	B uint8
	A uint8
}

type ColorFloat struct {
	R float32
	G float32
	B float32
	A float32
}

func NewColor(r uint8, g uint8, b uint8, a uint8) Color {
	return Color{r, g, b, a}
}

func NewColorFloat(r float32, g float32, b float32, a float32) ColorFloat {
	return ColorFloat{r, g, b, a}
}

func (color Color) ToFloat() ColorFloat {
	return NewColorFloat(float32(color.R)/255, float32(color.G)/255, float32(color.B)/255, float32(color.A)/255)
}

func (color ColorFloat) ToUint() Color {
	return NewColor(uint8(color.R)*255, uint8(color.G)*255, uint8(color.B)*255, uint8(color.A)*255)
}

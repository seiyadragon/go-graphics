package graphics

type Material struct {
	Shader   Shader
	Textures []Texture
}

func NewMaterial(shader Shader) Material {
	return Material{shader, nil}
}

func (m Material) Bind() {
	m.Shader.Bind()

	if m.Textures != nil {
		for i := 0; i < len(m.Textures); i++ {
			m.Textures[i].Bind(uint32(i))
		}
	}
}

func (m Material) Unbind() {
	m.Shader.Unbind()

	if m.Textures != nil {
		for i := 0; i < len(m.Textures); i++ {
			m.Textures[i].Unbind(uint32(i))
		}
	}
}

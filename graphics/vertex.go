package graphics

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

type Vertex struct {
	Position    mgl32.Vec3
	Texture     mgl32.Vec3
	Normals     mgl32.Vec3
	BoneIds     mgl32.Vec3
	BoneWeights mgl32.Vec3
}

type VAO struct {
	Id uint32
}

type VBO struct {
	Id uint32
}

type IBO struct {
	Id uint32
}

func NewVertex(position mgl32.Vec3, texture mgl32.Vec3, normals mgl32.Vec3) Vertex {
	return Vertex{position, texture, normals, mgl32.Vec3{0, 0, 0}, mgl32.Vec3{0, 0, 0}}
}

func NewVAO() VAO {
	var tmp uint32
	gl.GenVertexArrays(1, &tmp)
	vao := VAO{tmp}
	return vao
}

func (v VAO) Bind() {
	gl.BindVertexArray(v.Id)
}

func (v VAO) Unbind() {
	gl.BindVertexArray(0)
}

func (v VAO) SetVertexAttrib(attrib uint32, offset int) {
	v.Bind()
	gl.VertexAttribPointerWithOffset(attrib, 3, gl.FLOAT, false, 15*4, uintptr(offset))
	gl.EnableVertexAttribArray(attrib)
	v.Unbind()
}

func NewVBO(data []Vertex) VBO {
	var tmp uint32
	gl.GenBuffers(1, &tmp)
	vbo := VBO{tmp}
	vbo.Bind()
	gl.BufferData(gl.ARRAY_BUFFER, len(data)*(15*4), gl.Ptr(data), gl.STATIC_DRAW)
	vbo.Unbind()

	return vbo
}

func (v VBO) Bind() {
	gl.BindBuffer(gl.ARRAY_BUFFER, v.Id)
}

func (v VBO) Unbind() {
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)
}

func NewIBO(data []uint32) IBO {
	var tmp uint32
	gl.GenBuffers(1, &tmp)
	ibo := IBO{tmp}
	ibo.Bind()
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(data)*4, gl.Ptr(data), gl.STATIC_DRAW)
	ibo.Unbind()

	return ibo
}

func (v IBO) Bind() {
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, v.Id)
}

func (v IBO) Unbind() {
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, 0)
}

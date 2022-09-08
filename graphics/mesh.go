package graphics

import (
	"fmt"
	"log"
	"os"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/udhos/gwob"
)

type Mesh struct {
	Vao   VAO
	Count int32
}

func NewMesh(vertices []Vertex, indices []uint32) Mesh {
	vao := NewVAO()
	vbo := NewVBO(vertices)
	ibo := NewIBO(indices)

	vao.Bind()
	vbo.Bind()
	ibo.Bind()

	vao.SetVertexAttrib(0, 0)
	vao.SetVertexAttrib(1, 12)
	vao.Unbind()

	return Mesh{vao, int32(len(indices))}
}

func (m Mesh) Draw(material Material) {
	material.Bind()
	m.Vao.Bind()
	gl.DrawElements(gl.TRIANGLES, m.Count, gl.UNSIGNED_INT, nil)
	m.Vao.Unbind()
	material.Unbind()
}

func NewMeshFromFile(path string) Mesh {
	options := &gwob.ObjParserOptions{
		LogStats: true,
		Logger:   func(msg string) { fmt.Fprintln(os.Stderr, msg) },
	}

	o, errObj := gwob.NewObjFromFile(path, options)
	if errObj != nil {
		log.Printf("obj: parse error input=%s: %v", path, errObj)
	}

	var vertices []Vertex
	var indices []uint32

	for i := 0; i < len(o.Coord); i += 3 {
		vertices = append(vertices, NewVertex(mgl32.Vec3{o.Coord[i], o.Coord[i+1], o.Coord[i+2]}))
	}

	for i := 0; i < len(o.Indices); i++ {
		indices = append(indices, uint32(o.Indices[i]))
	}

	return NewMesh(vertices, indices)
}

func NewTriangle() Mesh {
	vertices := []Vertex{
		NewVertex(mgl32.Vec3{0, 1, 0}, mgl32.Vec3{0.5, 0, 0}),
		NewVertex(mgl32.Vec3{-1, -1, 0}, mgl32.Vec3{0, 1, 0}),
		NewVertex(mgl32.Vec3{1, -1, 0}, mgl32.Vec3{1, 1, 0}),
	}

	indices := []uint32{0, 1, 2}

	return NewMesh(vertices, indices)
}

func NewPlane() Mesh {
	vertices := []Vertex{
		NewVertex(mgl32.Vec3{-1, 1, 0}, mgl32.Vec3{0, 0, 0}),
		NewVertex(mgl32.Vec3{-1, -1, 0}, mgl32.Vec3{0, 1, 0}),
		NewVertex(mgl32.Vec3{1, -1, 0}, mgl32.Vec3{1, 1, 0}),
		NewVertex(mgl32.Vec3{1, 1, 0}, mgl32.Vec3{1, 0, 0}),
	}

	indices := []uint32{0, 1, 2, 2, 3, 0}

	return NewMesh(vertices, indices)
}

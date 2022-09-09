package graphics

import (
	"fmt"
	"os"
	"strings"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

type Shader struct {
	Program uint32
}

func NewShader(vertexSrc string, fragmentSrc string) Shader {
	program := gl.CreateProgram()
	vertex, err := compileShader(vertexSrc, gl.VERTEX_SHADER)
	if err != nil {
		panic(err)
	}
	fragment, err := compileShader(fragmentSrc, gl.FRAGMENT_SHADER)
	if err != nil {
		panic(err)
	}
	gl.UseProgram(program)

	gl.AttachShader(program, vertex)
	gl.AttachShader(program, fragment)
	gl.LinkProgram(program)

	gl.DeleteShader(vertex)
	gl.DeleteShader(fragment)

	gl.UseProgram(0)

	return Shader{program}
}

func NewShaderFromFile(vertex string, fragment string) Shader {
	vertexSrc, err := os.ReadFile(vertex)
	if err != nil {
		panic(err)
	}

	fragmentSrc, err := os.ReadFile(fragment)
	if err != nil {
		panic(err)
	}

	return NewShader(string(vertexSrc), string(fragmentSrc))
}

func (s Shader) Bind() {
	gl.UseProgram(s.Program)
}

func (s Shader) Unbind() {
	gl.UseProgram(0)
}

func (s Shader) SetUniform1f(name string, val float32) {
	s.Bind()
	gl.Uniform1f(gl.GetUniformLocation(s.Program, gl.Str(name+string(rune(0)))), val)
	s.Unbind()
}

func (s Shader) SetUniform1i(name string, val int32) {
	s.Bind()
	gl.Uniform1i(gl.GetUniformLocation(s.Program, gl.Str(name+string(rune(0)))), val)
	s.Unbind()
}

func (s Shader) SetUniform2f(name string, val mgl32.Vec2) {
	s.Bind()
	gl.Uniform2f(gl.GetUniformLocation(s.Program, gl.Str(name+string(rune(0)))), val.X(), val.Y())
	s.Unbind()
}

func (s Shader) SetUniform3f(name string, val mgl32.Vec3) {
	s.Bind()
	gl.Uniform3f(gl.GetUniformLocation(s.Program, gl.Str(name+string(rune(0)))), val.X(), val.Y(), val.Z())
	s.Unbind()
}

func (s Shader) SetUniform4f(name string, val mgl32.Vec4) {
	s.Bind()
	gl.Uniform4f(gl.GetUniformLocation(s.Program, gl.Str(name+string(rune(0)))), val.X(), val.Y(), val.Z(), val.W())
	s.Unbind()
}

func (s Shader) SetUniformMat4(name string, val mgl32.Mat4) {
	s.Bind()
	gl.UniformMatrix4fv(gl.GetUniformLocation(s.Program, gl.Str(name+string(rune(0)))), 1, false, &val[0])
	s.Unbind()
}

func compileShader(source string, shaderType uint32) (uint32, error) {
	shader := gl.CreateShader(shaderType)

	csources, free := gl.Strs(source)
	gl.ShaderSource(shader, 1, csources, nil)
	free()
	gl.CompileShader(shader)

	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(log))

		return 0, fmt.Errorf("failed to compile %v: %v", source, log)
	}

	return shader, nil
}

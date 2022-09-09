package graphics

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

func Init() (*glfw.Window, error) {
	err := glfw.Init()
	if err != nil {
		return nil, err
	}

	window, err := glfw.CreateWindow(1280, 720, "Window", nil, nil)
	if err != nil {
		return nil, err
	}

	window.MakeContextCurrent()

	if err := gl.Init(); err != nil {
		return nil, err
	}

	gl.Enable(gl.CULL_FACE)

	gl.Enable(gl.BLEND)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)

	gl.Enable(gl.DEPTH_TEST)
	gl.DepthFunc(gl.LESS)

	return window, err
}

func ClearScreen(color Color) {
	flt := color.ToVec4()
	gl.ClearColor(flt.X(), flt.Y(), flt.Z(), flt.W())
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
}

func UpdateWindow(window *glfw.Window) {
	window.SwapBuffers()
	glfw.PollEvents()
	width, height := window.GetSize()
	gl.Viewport(0, 0, int32(width), int32(height))
}

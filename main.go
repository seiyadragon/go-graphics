package main

import (
	_ "image/jpeg"
	_ "image/png"
	"runtime"

	"github.com/go-gl/mathgl/mgl32"
	"github.com/seiyadragon/go-graphics/graphics"
)

func init() {
	runtime.LockOSThread()
}

func main() {
	window, _ := graphics.Init()

	layer := graphics.DefaultLayer(graphics.NewView(1280, 720, mgl32.Vec3{0, 0, 6}, mgl32.Vec3{0, 0, -1}, mgl32.Vec3{0, 1, 0}))

	model := graphics.NewModel(
		graphics.NewTriangle(),
		graphics.NewMaterial(graphics.NewShaderFromFile("graphics/shaders/vertex.glsl", "graphics/shaders/fragment.glsl")),
		mgl32.Vec3{0, 0, 0},
		mgl32.Vec3{2, 3, 1},
		mgl32.Vec3{2, 2, 1},
	)

	model.Material.Shader.SetUniform4f("color", mgl32.Vec4{1, 0, 0, 1.0})
	model.Material.Shader.SetUniform1i("sampler_obj", 0)

	model.Material.AddTexture(graphics.NewTextureFromFile("test.jpg"))

	for !window.ShouldClose() {
		graphics.ClearScreen(graphics.NewColor(105, 15, 0, 255))

		layer.Draw(model)

		graphics.UpdateWindow(window)
	}
}

package main

import (
	_ "image/jpeg"
	_ "image/png"
	"math"
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
		graphics.NewMeshFromFile("cube.obj"),
		graphics.NewMaterial(graphics.NewShaderFromFile("graphics/shaders/vertex.glsl", "graphics/shaders/fragment.glsl")),
		mgl32.Vec3{0, 0, 10},
		mgl32.Vec3{2, 3, 1},
		mgl32.Vec3{2, 2, 1},
	)

	model.Material.Shader.SetUniform4f("color", graphics.NewColor(169, 42, 69, 255).ToVec4())
	model.Material.Shader.SetUniform1i("sampler_obj", 0)

	model.Material.AddTexture(graphics.NewTextureFromFile("test.jpg"))

	counter := 0.0

	for !window.ShouldClose() {
		graphics.ClearScreen(graphics.NewColor(35, 15, 115, 255))

		model.Rotation = model.Rotation.Add(mgl32.Vec3{float32(counter), float32(counter), float32(counter)})
		model.Position = mgl32.Vec3{float32(math.Sin(counter*1000) * 6), float32(math.Cos(counter*1000) * 3), -10}

		layer.Draw(model)
		counter += 0.00001

		graphics.UpdateWindow(window)
	}
}

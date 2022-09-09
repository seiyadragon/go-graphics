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

	var models []graphics.Model

	for i := 0; i < 16; i++ {
		models = append(models, graphics.NewModel(
			graphics.NewMeshFromFile("cube.obj"),
			graphics.NewMaterial(graphics.NewShaderFromFile("graphics/shaders/vertex.glsl", "graphics/shaders/fragment.glsl")),
			mgl32.Vec3{0, 0, 10},
			mgl32.Vec3{2, 3, 1},
			mgl32.Vec3{2, 2, 1},
		))

		models[i].Material.Shader.SetUniform4f("color", graphics.NewColor(169, 42, 69, 255).ToVec4())
		models[i].Material.Shader.SetUniform1i("sampler_obj", 1)
		models[i].Material.Shader.SetUniform1i("sampler_obj2", 0)
		models[i].Material.Textures = append(models[i].Material.Textures, graphics.NewTextureFromFile("test.jpg"))
		models[i].Material.Textures = append(models[i].Material.Textures, graphics.NewTextureFromFile("test2.jpg"))
	}

	counter := 0.0

	for !window.ShouldClose() {
		graphics.ClearScreen(graphics.NewColor(35, 15, 115, 255))

		for i := 0; i < 16; i++ {
			models[i].Rotation = models[i].Rotation.Add(mgl32.Vec3{float32(counter), float32(counter), float32(counter)})
			models[i].Position = mgl32.Vec3{float32(math.Tan(float64(i))*8) + float32(math.Sin(counter*1000)*6),
				float32(math.Sin(float64(i))*8) + float32(math.Cos(counter*1000)*3),
				-20 - float32(math.Cos(float64(i))*8) + float32(math.Tan(counter*1000)*6)}

			layer.Draw(models[i])
		}

		counter += 0.00001

		graphics.UpdateWindow(window)
	}
}

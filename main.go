package main

import (
	_ "image/jpeg"
	_ "image/png"
	"math"
	"runtime"
	"time"

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
	ball_amount := 25
	ball_model := graphics.NewMeshFromFile("res/cube.obj")
	shader := graphics.NewShaderFromFile("res/shader.glsl")
	tex1 := graphics.NewTextureFromFile("res/test.jpg")
	tex2 := graphics.NewTextureFromFile("res/test2.jpg")

	for i := 0; i < ball_amount; i++ {
		models = append(models, graphics.NewModel(
			ball_model,
			graphics.NewMaterial(shader),
			mgl32.Vec3{0, 0, 10},
			mgl32.Vec3{2, 3, 1},
			mgl32.Vec3{2, 2, 1},
		))

		models[i].Material.Shader.SetUniform4f("color", graphics.NewColor(169, 42, 69, 255).ToVec4())
		models[i].Material.Shader.SetUniform1i("sampler_obj", 1)
		models[i].Material.Shader.SetUniform1i("sampler_obj2", 0)
		models[i].Material.Textures = append(models[i].Material.Textures, tex1)
		models[i].Material.Textures = append(models[i].Material.Textures, tex2)
	}

	counter := 0.0
	lastTime := time.Now()
	fps := 0

	for !window.ShouldClose() {
		now := time.Now()
		fps++

		if now.Unix()-lastTime.Unix() >= 1 {
			println(fps)
			lastTime = time.Now()
			fps = 0
		}

		graphics.ClearScreen(graphics.NewColor(35, 15, 115, 255))

		for i := 0; i < ball_amount; i++ {
			models[i].Rotation = models[i].Rotation.Add(mgl32.Vec3{float32(counter), float32(counter), float32(counter)})
			models[i].Position = mgl32.Vec3{float32(math.Tan(float64(i))*8) + float32(math.Sin(counter*1000)*6),
				float32(math.Sin(float64(i))*8) + float32(math.Cos(counter*1000)*3),
				-20 - float32(math.Cos(float64(i))*8) + float32(math.Tan(counter*1000)*6)}

			layer.DrawModel(models[i])
		}

		counter += 0.00001

		graphics.UpdateWindow(window)
	}
}

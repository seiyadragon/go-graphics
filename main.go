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
	strip := graphics.NewTextureFromFile("res/strip2.png")

	var frames []graphics.Texture

	for i := 0; i < int(22); i++ {
		subimage := strip.GetSubTexture(int32(i*144), 0, 144, 64)

		frames = append(frames, subimage)
	}

	for i := 0; i < ball_amount; i++ {
		models = append(models, graphics.NewModel(
			ball_model,
			graphics.NewMaterial(shader),
			mgl32.Vec3{0, 0, 10},
			mgl32.Vec3{2, 3, 1},
			mgl32.Vec3{0.5, 0.5, 0.5},
		))

		models[i].Material.Shader.SetUniform4f("color", graphics.NewColor(169, 42, 69, 255).ToVec4())
		models[i].Material.Shader.SetUniform1i("sampler_obj", 0)
		models[i].Material.Textures = append(models[i].Material.Textures, tex1)
	}

	sprite := graphics.NewSprite(graphics.NewMaterial(shader), mgl32.Vec3{0, 0, -10}, mgl32.Vec3{0, 0, 0}, mgl32.Vec3{5, 5, 1})
	sprite.Material.Shader.SetUniform4f("color", graphics.NewColor(255, 255, 255, 255).ToVec4())
	sprite.Material.Shader.SetUniform1i("sampler_obj", 0)

	animation := graphics.NewAnimation2D(60, frames)

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

			layer.Draw(models[i])
		}

		sprite.Material.Textures = append(sprite.Material.Textures, animation.GetAnimationFrame())
		layer.Draw(sprite)
		sprite.Material.Textures = append(sprite.Material.Textures[:0], sprite.Material.Textures[0+1:]...)

		counter += 0.00001

		graphics.UpdateWindow(window)
	}
}

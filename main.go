package main

import (
	_ "image/jpeg"
	_ "image/png"
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
	monkey := graphics.NewModel(
		graphics.NewMeshFromFile("res/monkey.obj"),
		graphics.NewMaterial(graphics.NewShaderFromFile("res/shader.glsl")),
		mgl32.Vec3{0, 0, -5},
		mgl32.Vec3{0, 0, 0},
		mgl32.Vec3{3, 3, 3},
	)

	monkey.Material.Shader.SetUniform1i("sampler", 0)
	monkey.Material.Textures = append(monkey.Material.Textures, graphics.NewTextureFromFile("res/fur.jpg"))

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

		layer.Draw(monkey)

		graphics.UpdateWindow(window)
	}
}

package graphics

import (
	"time"

	"github.com/go-gl/mathgl/mgl32"
)

type Animation2D struct {
	FPS            int
	Frames         []Texture
	LastFrameIndex int
	LastFrameTime  int64
}

func NewAnimation2D(fps int, frames []Texture) Animation2D {
	return Animation2D{fps, frames, 0, time.Now().UnixMilli()}
}

func (a *Animation2D) GetAnimationFrame() Texture {
	now := time.Now().UnixMilli()

	if now-a.LastFrameTime >= int64(1000/a.FPS) {
		a.LastFrameTime = time.Now().UnixMilli()
		a.LastFrameIndex++

		if a.LastFrameIndex > len(a.Frames)-1 {
			a.LastFrameIndex = 0
		}
	}

	return a.Frames[a.LastFrameIndex]
}

type Joint struct {
	Id               float32
	Children         []Joint
	Transform        mgl32.Mat4
	InverseTransform mgl32.Mat4
}

func NewJoint(id float32) Joint {
	joint := Joint{id, nil, mgl32.Ident4(), mgl32.Ident4()}

	return joint
}

func (j Joint) CalculateInveseTransform(parentTransform mgl32.Mat4) {
	trans := parentTransform.Mul4(j.Transform)
	j.InverseTransform = trans.Inv()

	for i := 0; i < len(j.Children); i++ {
		j.Children[i].CalculateInveseTransform(trans)
	}
}

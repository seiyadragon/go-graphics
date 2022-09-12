package graphics

import "time"

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

package graphics

import (
	"image"
	"os"

	"github.com/go-gl/gl/v4.1-core/gl"
)

type Texture struct {
	Id uint32
}

func NewTexture(data image.Image) Texture {
	var tex uint32
	var pixels []uint8

	width, height := int32(data.Bounds().Max.X), int32(data.Bounds().Max.Y)

	for y := 0; y < int(height); y++ {
		for x := 0; x < int(width); x++ {
			r, g, b, a := data.At(x, y).RGBA()
			pixels = append(pixels, uint8(r/257))
			pixels = append(pixels, uint8(g/257))
			pixels = append(pixels, uint8(b/257))
			pixels = append(pixels, uint8(a/257))
		}
	}

	gl.GenTextures(1, &tex)
	gl.BindTexture(gl.TEXTURE_2D, tex)

	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_BORDER)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_BORDER)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR_MIPMAP_LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR_MIPMAP_LINEAR)

	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA, width, height, 0, gl.RGBA, gl.UNSIGNED_BYTE, gl.Ptr(pixels))
	gl.GenerateMipmap(gl.TEXTURE_2D)

	gl.BindTexture(gl.TEXTURE_2D, 0)

	return Texture{tex}
}

func NewTextureFromFile(path string) Texture {
	image, err := LoadImageFromFile(path)
	if err != nil {
		panic(err)
	}

	return NewTexture(image)
}

func (t Texture) Bind(unit uint32) {
	gl.ActiveTexture(gl.TEXTURE0 + unit)
	gl.BindTexture(gl.TEXTURE_2D, t.Id)
}

func (t Texture) Unbind(unit uint32) {
	gl.ActiveTexture(gl.TEXTURE0 + unit)
	gl.BindTexture(gl.TEXTURE_2D, 0)
}

func (t Texture) GetSubTexture(x, y, width, height int32) Texture {
	subimage := make([]uint8, width*height*4)

	t.Bind(0)
	gl.GetTextureSubImage(t.Id, 0, x, y, 0, width, height, 1, gl.RGBA, gl.UNSIGNED_BYTE, int32(len(subimage)*2), gl.Ptr(subimage))
	t.Unbind(0)

	var tex uint32

	gl.GenTextures(1, &tex)
	gl.BindTexture(gl.TEXTURE_2D, tex)

	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_BORDER)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_BORDER)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR_MIPMAP_LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR_MIPMAP_LINEAR)

	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA, width, height, 0, gl.RGBA, gl.UNSIGNED_BYTE, gl.Ptr(subimage))
	gl.GenerateMipmap(gl.TEXTURE_2D)

	gl.BindTexture(gl.TEXTURE_2D, 0)

	return Texture{tex}
}

func LoadImageFromFile(filePath string) (image.Image, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	image, _, err := image.Decode(f)
	return image, err
}

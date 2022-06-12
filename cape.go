package vails

import (
	"github.com/df-mc/dragonfly/server/player/skin"
	"image/png"
	"os"
)

// Cape stores data about a cosmetic cape, such as it's name, image, etc.
type Cape struct {
	name string
	plus bool
	cape skin.Cape
}

// NewCape creates a new cape with the given name and image. The plus parameter determines whether the cape can only
// be accessed by players with the Plus role.
func NewCape(name string, path string, plus bool) Cape {
	return Cape{
		name: name,
		plus: plus,
		cape: read("assets/capes/" + path),
	}
}

// Name returns the name of the cape.
func (c Cape) Name() string {
	return c.name
}

// Premium returns true if the cape can only be accessed by players with the Plus role.
func (c Cape) Premium() bool {
	return c.plus
}

// Cape returns the image data of the cape.
func (c Cape) Cape() skin.Cape {
	return c.cape
}

// read performs a read on the path provided and returns a dragonfly cape.
func read(path string) skin.Cape {
	f, _ := os.OpenFile(path, os.O_RDONLY, os.ModePerm)
	defer f.Close()
	i, err := png.Decode(f)
	if err != nil {
		panic(err)
	}

	c := skin.NewCape(i.Bounds().Max.X, i.Bounds().Max.Y)
	for y := 0; y < i.Bounds().Max.Y; y++ {
		for x := 0; x < i.Bounds().Max.X; x++ {
			color := i.At(x, y)
			r, g, b, a := color.RGBA()
			i := x*4 + i.Bounds().Max.X*y*4
			c.Pix[i], c.Pix[i+1], c.Pix[i+2], c.Pix[i+3] = uint8(r), uint8(g), uint8(b), uint8(a)
		}
	}
	return c
}

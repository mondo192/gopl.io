package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
)

var (
	red = color.RGBA{0xff, 0x00, 0x00, 0xff}
	green = color.RGBA{0x00, 0xff, 0x00, 0xff}
	blue = color.RGBA{0x00, 0x00, 0xff, 0xff}
)

var palette = []color.Color{color.Black, red, red, green, green}

const (
	whiteindex = 0
	blackindex = 1
)

func lissajous(out io.Writer) {
	const (
		cycles  = 5
		res     = 0.001
		size    = 100
		nframes = 64
		delay   = 8
	)
	frequency := rand.Float64() * 3.0
	animation := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rectangle := image.Rect(0, 0, 2*size+1, 2*size+1)
		image := image.NewPaletted(rectangle, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*frequency + phase)
			colorIndex := uint8(i%(len(palette)-1) + 1)
			image.SetColorIndex(size+int(x*size+0.5), int(y*size+0.5), colorIndex)
		}
		phase += 0.1
		animation.Delay = append(animation.Delay, delay)
		animation.Image = append(animation.Image, image)
	}
	gif.EncodeAll(out, &animation)
}

func main() {
	lissajous(os.Stdout)
}

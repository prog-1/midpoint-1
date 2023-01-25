package main

import (
	"fmt"
	"image/color"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	screenWidth  = 640
	screenHeight = 480
)

type Line struct {
	X1, Y1    int
	Magnitude int
	Degrees   float64
	color.Color
}

func ToRadians(Degrees float64) float64 {
	return Degrees * math.Pi / float64(180)
}

type game struct {
	l Line
}

func (g *game) Layout(outWidth, outHeight int) (w, h int) { return screenWidth, screenHeight }
func (g *game) Update() error {
	g.l.Degrees += 1
	return nil
}
func (g *game) Draw(screen *ebiten.Image) {
	x := float64(g.l.Magnitude) * math.Cos(ToRadians(g.l.Degrees))
	y := float64(g.l.Magnitude) * math.Sin(ToRadians(g.l.Degrees))
	x2, y2 := x+float64(g.l.X1), y+float64(g.l.Y1)

	DrawLineDDA(screen, g.l.X1, g.l.Y1, int(x2), int(y2), g.l.Color)
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	g := game{Line{320, 240, 100, 0, color.RGBA{255, 1, 1, 255}}}
	if err := ebiten.RunGame(&g); err != nil {
		log.Fatal(err)
	}
}

// DrawLineDDA rasterizes a line using Digital Differential Analyzer algorithm.
func DrawLineDDA(img *ebiten.Image, x1, y1, x2, y2 int, c color.Color) {
	if x2 < x1 {
		x1, x2 = x2, x1
		y1, y2 = y2, y1
	}
	Dx := x2 - x1
	Dy := y2 - y1
	A := float64(Dy)
	B := float64(-Dx)
	// C = b = y - kx = y1 - kx1
	k := Dy / Dx
	C := float64(y1 - k + x1)

	// xm, ym - middle pointn coord.
	xm, ym := float64(x1+1), float64(y1)+0.5
	d1 := A*xm + B*ym + C
	ebitenutil.DebugPrint(img, fmt.Sprint(d1))
	if d1 >= 0 {
		img.Set(x1+1, y1+1, c)
	}
	for x := x1; x < x2; x++ {

	}
}

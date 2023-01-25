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
	if math.Abs(g.l.Degrees) > 45 {
		g.l.Degrees = 0
	}
	return nil
}
func (g *game) Draw(screen *ebiten.Image) {
	DrawLine(screen, g.l.X1, g.l.Y1, 10, 20, g.l.Color)
	ebitenutil.DebugPrint(screen, fmt.Sprint(d))
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	g := game{Line{320, 240, 100, 0, color.RGBA{255, 255, 255, 255}}}
	if err := ebiten.RunGame(&g); err != nil {
		log.Fatal(err)
	}
}

var d float64

func DrawLine(img *ebiten.Image, x1, y1, x2, y2 int, c color.Color) {
	if x2 < x1 {
		x1, x2 = x2, x1
		y1, y2 = y2, y1
	}
	Dx, Dy := x2-x1, y2-y1
	A, B, C := Dy, -Dx, Dx*y1-Dy*x1

	f := func(x, y float64) float64 {
		return float64(A)*x + float64(B)*y + float64(C)
	}

	img.Set(x2, y2, c)

	for x, y := x1, y1; x < x2; x++ {
		img.Set(x, y, c)
		xm, ym := x+1, float64(y)-0.5
		d = f(float64(xm), ym)
		if d < 0 {
			y--
		}
	}
}

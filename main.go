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

type game struct{}

func (g *game) Layout(outWidth, outHeight int) (w, h int) { return screenWidth, screenHeight }
func (g *game) Update() error                             { return nil }
func (g *game) Draw(screen *ebiten.Image) {
	DrawLine(screen, 320, 240, 400, 100, color.White)
	ebitenutil.DebugPrint(screen, fmt.Sprint(d))
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	if err := ebiten.RunGame(&game{}); err != nil {
		log.Fatal(err)
	}
}

var d float64

func DrawLine(img *ebiten.Image, x1, y1, x2, y2 int, c color.Color) {
	img.Set(x1, y1, color.RGBA{1, 255, 1, 255})
	img.Set(x2, y2, color.RGBA{1, 255, 1, 255})

	Dx, Dy := x2-x1, y2-y1
	k := float64(Dy) / float64(Dx)
	if math.Abs(k) <= 1 {
		for x, y := x1, y1; x < x2; x++ {
			img.Set(x, y, c)
			xm, ym := float64(x+1), float64(y)-0.5
			f := func(x, y float64) float64 {
				A, B, C := Dy, -Dx, Dx*y1-Dy*x1
				return float64(A)*x + float64(B)*y + float64(C)
			}
			d = f(xm, ym)
			if d < 0 {
				y--
			}
		}
	} else {
		for x, y := x1, y1; y > y2; y-- {
			img.Set(x, y, c)
			xm, ym := float64(x)+0.5, float64(y-1)
			f := func(x, y float64) float64 {
				A, B, C := -Dy, Dx, Dx*x1-Dy*y1
				return float64(A)*y + float64(B)*x + float64(C)
			}
			d = f(xm, ym)
			if d < 0 {
				x++
			}
		}
	}

}

package main

import (
	"image/color"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type (
	point struct {
		x, y float64
	}

	Game struct {
		p1, p2 point
	}
)

var col = color.RGBA{R: 0xff, G: 0xff, B: 0xff, A: 0xff}

const (
	winTitle            = "raster"
	winWidth, winHeight = 500, 500
)

func DrawLineDDA(img *ebiten.Image, x1, y1, x2, y2 float64, col color.Color) {
	if x2 < x1 {
		x1, x2 = x2, x1
		y1, y2 = y2, y1
	}
	dx, dy := x2-x1, y2-y1

	if math.Abs(dx-dy) > 0 {
		a, b, c := dy, -dx, dx*y1-dy*x1
		dirY := 1.0
		if dy < 0 {
			dirY = -1
		}
		for x, y := x1, y1; x <= x2; x++ {
			d := a*(x) + b*(y+0.5) + c
			if d*dirY > 0 {
				y += dirY
			}
			img.Set(int(x), int(y), col)
		}

	} else {
		a, b, c := dx, -dy, dy*x1-dx*y1
		dirX := 1.0
		if dx < 0 {
			dirX = -1
		}
		for x, y := x1, y1; y <= y2; y++ {
			d := a*(y) + b*(x+0.5) + c
			if d*dirX > 0 {
				y += dirX
			}
			img.Set(int(x), int(y), col)
		}

	}

}
func (g *Game) Draw(screen *ebiten.Image) {
	DrawLineDDA(screen, g.p1.x, g.p1.y, g.p2.x, g.p2.y, col)

}

func (g *Game) rotation() {
	g.p1.x = g.p1.x*math.Cos(0.03) - g.p1.y*math.Sin(0.03) + (250 - 250*math.Cos(0.03) + 250*math.Sin(0.03))
	g.p1.y = g.p1.x*math.Sin(0.03) + g.p1.y*math.Cos(0.03) + (-250*math.Sin(0.03) + 250 - 250*math.Cos(0.03))
	g.p2.x = g.p2.x*math.Cos(0.03) - g.p2.y*math.Sin(0.03) + (250 - 250*math.Cos(0.03) + 250*math.Sin(0.03))
	g.p2.y = g.p2.x*math.Sin(0.03) + g.p2.y*math.Cos(0.03) + (-250*math.Sin(0.03) + 250 - 250*math.Cos(0.03))
}
func (g *Game) Update() error {
	g.rotation()
	return nil

}

func (g *Game) Layout(int, int) (w, h int) { return winWidth, winHeight }

func main() {
	ebiten.SetWindowTitle(winTitle)
	ebiten.SetWindowSize(winWidth, winHeight)
	ebiten.SetWindowResizable(true)
	if err := ebiten.RunGame(&Game{p1: point{x: 100, y: 100}, p2: point{x: 400, y: 400}}); err != nil {
		log.Fatal(err)
	}
}

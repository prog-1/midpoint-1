package main

import (
	"image/color"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 640
	screenHeight = 480
)

type Line struct {
	x1, y1  int
	length  int
	radians float64
	c       color.Color
}

type game struct {
	l *Line
}

func (g *game) Layout(outWidth, outHeight int) (w, h int) { return screenWidth, screenHeight }
func (g *game) Update() error {
	g.l.radians += math.Pi / 180
	return nil
}
func (g *game) Draw(screen *ebiten.Image) {
	x := float64(g.l.length) * math.Cos(g.l.radians)
	y := float64(g.l.length) * math.Sin(g.l.radians)
	x2, y2 := g.l.x1+int(x), g.l.y1+int(y)
	DrawLine(screen, g.l.x1, g.l.y1, x2, y2, g.l.c)
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	g := game{&Line{screenWidth / 2, screenHeight / 2, 100, 0, color.RGBA{1, 100, 100, 255}}}
	if err := ebiten.RunGame(&g); err != nil {
		log.Fatal(err)
	}
}

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}
func DrawLine(img *ebiten.Image, x1, y1, x2, y2 int, c color.Color) {
	// abs(Dy) < abs(dx) | / abs(dx) => abs(Dy)/abs(Dx) < 1 == abs(k) < 1
	if abs(y2-y1) < abs(x2-x1) {
		if x1 > x2 {
			x1, x2 = x2, x1
			y1, y2 = y2, y1
		}
		dx, dy := x2-x1, y2-y1
		dirY := 1
		// Dy < 0 => y2 - y1 < 0 => y1 > y2 => Growing downwards
		if dy < 0 {
			dirY = -1
			dy = -dy // For us to pretend that line is growing upwards
		}
		d := 2*dy - dx
		for x, y := x1, y1; x < x2; x++ {
			img.Set(x, y, c)
			if d >= 0 { // NE
				y += dirY
				d += dy - dx
			} else { // E
				d += dy
			}
		}
	} else {
		if y1 > y2 {
			x1, x2 = x2, x1
			y1, y2 = y2, y1
		}
		dx, dy := x2-x1, y2-y1
		dirX := 1
		if dx < 0 {
			dirX = -1
			dx = -dx
		}
		d := 2*dx - dy
		for x, y := x1, y1; y < y2; y++ {
			img.Set(x, y, c)
			if d >= 0 { // NE
				x += dirX
				d += dx - dy
			} else { // E
				d += dx
			}
		}
	}

	// img.Set(x1, y1, color.RGBA{1, 255, 1, 255})
	// img.Set(x2, y2, color.RGBA{1, 255, 1, 255})
}

package main

import (
	"image/color"
	"log"
	"math"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten"
)

const (
	screenWidth  = 640
	screenHeight = 480
)

type Game struct {
	width, height int
}

func DrawLine(img *ebiten.Image, x1, y1, x2, y2 float64, color color.Color) {
	if math.Abs(x2-x1) >= math.Abs(y2-y1) {
		if x2 < x1 {
			x1, x2 = x2, x1
			y1, y2 = y2, y1
		}
		a, b := y2-y1, x1-x2
		dirY := 1.0
		if a < 0 {
			dirY = -1
			a = -a
		}
		d := a + b/2
		for x, y := x1, y1; x <= x2; x = x + 1 {
			img.Set(int(x), int(y), color)
			if d > 0 {
				d += b
				y += dirY
			}
			d += a
		}
	} else {
		if y2 < y1 {
			x1, x2 = x2, x1
			y1, y2 = y2, y1
		}
		a, b := x2-x1, y1-y2
		dirX := 1.0
		if a < 0 {
			dirX = -1
			a = -a
		}
		d := a + b/2
		for x, y := x1, y1; y <= y2; y = y + 1 {
			img.Set(int(x), int(y), color)
			if d > 0 {
				d += b
				x += dirX
			}
			d += a
		}
	}
}

func NewGame(width, height int) *Game {
	return &Game{
		width:  width,
		height: height,
	}
}

func (g *Game) Layout(outWidth, outHeight int) (w, h int) {
	return g.width, g.height
}

func (g *Game) Update(screen *ebiten.Image) error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	DrawLine(screen, 320, 240, 500, 100, color.White)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	g := NewGame(screenWidth, screenHeight)
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}

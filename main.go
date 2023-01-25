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

func DrawLine(img *ebiten.Image, x1, y1, x2, y2 int, color color.Color) {
	if math.Abs(float64(x2-x1)) >= math.Abs(float64(y2-y1)) {
		if x2 < x1 {
			x1, x2 = x2, x1
			y1, y2 = y2, y1
		}
		a, b, c := float64(y2-y1), float64(-(x2 - x1)), float64(x2*y1-y2*x1)
		if y1 > y2 {
			for x, y := float64(x1), float64(y1); int(x) <= x2; x = x + 1 {
				if a*(x+1)+b*(y-0.5)+c <= 0 {
					img.Set(int(x)+1, int(y)-1, color)
					y--
				} else {
					img.Set(int(x)+1, int(y), color)
				}
			}
		} else {
			for x, y := float64(x1), float64(y1); int(x) <= x2; x = x + 1 {
				if a*(x+1)+b*(y+0.5)+c <= 0 {
					img.Set(int(x)+1, int(y), color)
				} else {
					img.Set(int(x)+1, int(y)+1, color)
					y++
				}
			}
		}
	} else {
		if y2 < y1 {
			x1, x2 = x2, x1
			y1, y2 = y2, y1
		}
		a, b, c := float64(y2-y1), float64(-(x2 - x1)), float64(x2*y1-y2*x1)
		if x1 > x2 {
			for x, y := float64(x1), float64(y1); int(y) <= y2; y = y + 1 {
				if a*(x-0.5)+b*(y+1)+c <= 0 {
					img.Set(int(x), int(y)+1, color)
				} else {
					img.Set(int(x)-1, int(y)+1, color)
					x--
				}
			}
		} else {
			for x, y := float64(x1), float64(y1); int(y) <= y2; y = y + 1 {
				if a*(x+0.5)+b*(y+1)+c <= 0 {
					img.Set(int(x)+1, int(y)+1, color)
					x++
				} else {
					img.Set(int(x), int(y)+1, color)
				}
			}
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
	DrawLine(screen, g.width/2, g.height/2, 500, 100, color.White)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	g := NewGame(screenWidth, screenHeight)
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"image/color"
	"log"
	"math"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type Point struct {
	x, y float64
}

const (
	screenWidth  = 640
	screenHeight = 480
)

func DrawLine(img *ebiten.Image, x1, y1, x2, y2 float64, c color.Color) {
	if math.Abs(x2-x1) >= math.Abs(y2-y1) {
		if x2 < x1 {
			x1, x2 = x2, x1
			y1, y2 = y2, y1
		}

		A := y2 - y1
		B := -(x2 - x1)
		C := (x2-x1)*y1 - (y2-y1)*x1

		//  \ +
		//   \
		//  - \
		//     \
		if A >= 0 && B >= 0 {
			for x, y := x1, y1; x <= x2; x += 1 {
				if A*(x+1)+B*(y-0.5)+C <= 0 {
					img.Set(int(x)+1, int(y)-1, c)
					y--
				} else {
					img.Set(int(x)+1, int(y), c)
				}
			}
		} else if A < 0 && B < 0 {
			//  \ -
			//   \
			//  + \
			//     \
			for x, y := x1, y1; x <= x2; x += 1 {
				if A*(x+1)+B*(y-0.5)+C <= 0 {
					img.Set(int(x)+1, int(y), c)
				} else {
					img.Set(int(x)+1, int(y)-1, c)
					y--
				}
			}
		} else if A > 0 && B < 0 {
			//   - /
			//    /
			//   / +
			//  /
			for x, y := x1, y1; x <= x2; x += 1 {
				if A*(x+1)+B*(y+0.5)+C <= 0 {
					img.Set(int(x)+1, int(y), c)
				} else {
					img.Set(int(x)+1, int(y)+1, c)
					y++
				}
			}
		} else if A < 0 && B > 0 {
			//   + /
			//    /
			//   / -
			//  /
			for x, y := x1, y1; x <= x2; x += 1 {
				if A*(x+1)+B*(y+0.5)+C <= 0 {
					img.Set(int(x)+1, int(y)+1, c)
					y++
				} else {
					img.Set(int(x)+1, int(y), c)
				}
			}
		}
	} else {
		if y2 < y1 {
			x1, x2 = x2, x1
			y1, y2 = y2, y1
		}
		A := y2 - y1
		B := -(x2 - x1)
		C := (x2-x1)*y1 - (y2-y1)*x1
		if x1 < x2 {
			for x, y := x1, y1; y <= y2; y += 1 {
				if A*(x+0.5)+B*(y+1)+C <= 0 {
					img.Set(int(x)+1, int(y)+1, c)
					x++
				} else {
					img.Set(int(x), int(y)+1, c)
				}
			}
		} else {
			for x, y := x1, y1; y <= y2; y += 1 {
				if A*(x-0.5)+B*(y+1)+C <= 0 {
					img.Set(int(x), int(y)+1, c)
				} else {
					img.Set(int(x)-1, int(y)+1, c)
					x--
				}
			}
		}
	}
}

type Game struct {
	d          float64
	pos1, pos2 Point
}

func (g *Game) Update() error {
	g.d++
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ang := math.Pi * g.d / 180
	g.pos2.x = 0.2 * (g.pos1.x*math.Cos(ang) + g.pos1.y*math.Sin(ang))
	g.pos2.y = 0.2 * (-g.pos1.x*math.Sin(ang) + g.pos1.y*math.Cos(ang))
	g.pos2.x += (g.pos1.x)
	g.pos2.y += (g.pos1.y)
	DrawLine(screen, g.pos1.x, g.pos1.y, g.pos2.x, g.pos2.y, color.RGBA{R: 227, G: 76, B: 235, A: 1})
}

func (g *Game) Layout(int, int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	rand.Seed(time.Now().UnixNano())
	if err := ebiten.RunGame(&Game{pos1: Point{x: 320, y: 240}, pos2: Point{x: 0, y: 0}}); err != nil {
		log.Fatal(err)
	}
}

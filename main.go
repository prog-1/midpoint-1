package main

import (
	"image/color"
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten"
)

const (
	screenWidth  = 1290
	screenHeight = 960
)

type Point struct {
	x, y float64
}

type Game struct {
	width, height int
	last          time.Time
	p1, p2        Point
}

func (g *Game) DrawLine(img *ebiten.Image, c color.Color) {
	curx, cury := g.p1.x, g.p1.y
	img.Set(int(curx), int(cury), color.RGBA{255, 1, 1, 255})

	if g.p2.x < g.p1.x {
		g.p1.x, g.p1.y, g.p2.x, g.p2.y = g.p2.x, g.p2.y, g.p1.x, g.p1.y
	}
	for int(curx) != int(g.p2.x) {
		curx += 1
		cury += 0.5
		if g.Solve(curx, cury) {
			cury += 0.5
			img.Set(int(curx), int(cury), c)
		} else {
			cury -= 0.5
			img.Set(int(curx), int(cury), c)
		}
	}
}

// ^y*x + (-^x*y)+ ^x*y1 - ^y*x1
func (g *Game) Solve(x, y float64) bool {
	delx, dely := (g.p2.x - g.p1.x), (g.p2.y - g.p1.y)
	if (dely*x)+(-delx*y)+(delx*g.p1.y)-(dely*g.p1.x) > 0 {
		return true
	}
	return false
}

func (g *Game) Layout(outWidth, outHeight int) (w, h int) {
	return g.width, g.height
}

func (g *Game) Update(screen *ebiten.Image) error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.DrawLine(screen, color.RGBA{255, 255, 255, 255})
}

func NewGame(width, height int) *Game {
	return &Game{
		width:  width,
		height: height,
		p1:     Point{float64(width / 2), float64(height / 2)},
		p2:     Point{float64(width/2 + 1), float64(height/2 + 100)},
		last:   time.Now(),
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	ebiten.SetWindowSize(screenWidth, screenHeight)
	g := NewGame(screenWidth, screenHeight)
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}

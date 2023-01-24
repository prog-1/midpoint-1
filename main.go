package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 640
	screenHeight = 480
)

// Point is a struct for representing 2D vectors.
type Point struct {
	x, y int
}

type Line struct {
	pos1, pos2 Point
	color      color.RGBA
}

// NewLine initializes and returns a new Line instance.
func NewLine(x1, y1, x2, y2 int) *Line {
	return &Line{
		pos1: Point{x: x1, y: y1},
		pos2: Point{x: x2, y: y2},
		color: color.RGBA{
			R: 0xff,
			G: 0xff,
			B: 0xff,
			A: 0xff,
		},
	}
}

func DrawLine(img *ebiten.Image, x1, y1, x2, y2 int, c color.Color) {
	if x2 < x1 {
		x1, x2 = x2, x1
		y1, y2 = y2, y1
	}
	x, y := x1, y1
	A := (float64(y2 - y1))
	B := -(float64(x2 - x1))
	C := -B*float64(y1) - A*float64(x1)
	dy := 1
	if y1 > y2 {
		dy = -1
	}
	for x != x2 && y != y2 {
		img.Set(x, y, c)
		f := A*float64(x+1) + B*(float64(y)+0.5) + C
		x += 1
		if f > 0 {
			y += dy
		}
	}
}

func (l *Line) Draw(screen *ebiten.Image) {
	DrawLine(screen, l.pos1.x, l.pos1.y, l.pos2.x, l.pos2.y, l.color)
	DrawLine(screen, 300, 100, 100, 420, l.color) // additional
	DrawLine(screen, 500, 100, 500, 450, l.color) // additional  ?
	DrawLine(screen, 50, 450, 500, 450, l.color)  // additional  ?
	DrawLine(screen, 80, 300, 110, 50, l.color)   // additional  ?
}

// Game is a game instance.
type Game struct {
	width, height int
	line          *Line
}

// NewGame returns a new Game instance.
func NewGame(width, height int) *Game {
	return &Game{
		width:  width,
		height: height,
		line:   NewLine(100, 100, 400, 400),
	}
}

func (g *Game) Layout(outWidth, outHeight int) (w, h int) {
	return g.width, g.height
}

// Update updates a game state.
func (g *Game) Update() error {
	return nil
}

// Draw renders a game screen.
func (g *Game) Draw(screen *ebiten.Image) {
	g.line.Draw(screen)
}

func main() {
	//rand.Seed(time.Now().UnixNano())
	ebiten.SetWindowSize(screenWidth, screenHeight)
	g := NewGame(screenWidth, screenHeight)
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}

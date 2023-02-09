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

// Point is a struct for representing 2D vectors.
type Point struct {
	x, y int
}

type Line struct {
	pos1, pos2     Point
	radians        float64
	startx, starty float64
	color          color.RGBA
}

// These 2 functions are needed so that the initial length of the line always remains the same

func StartingXPos(x1, y1, x2, y2 int, radians float64) float64 {
	if x1 == x2 { // if the line is vertical
		return float64(y2-y1) / math.Sin(radians)
	}
	return float64(x2-x1) / math.Cos(radians)
}

func StartingYPos(x1, y1, x2, y2 int, radians float64) float64 {
	if y1 == y2 { // if the line is horizontal
		return float64(x2-x1) / math.Cos(radians)
	}
	return float64(y2-y1) / math.Sin(radians)
}

// NewLine initializes and returns a new Line instance.
func NewLine(x1, y1, x2, y2 int) *Line {
	return &Line{
		pos1:    Point{x: x1, y: y1},
		pos2:    Point{x: x2, y: y2},
		radians: math.Atan(float64(y2-y1) / float64(x2-x1)), // math.Atan(k)  k выражено из уравнения прямой, проходящей через две точки
		startx:  StartingXPos(x1, y1, x2, y2, math.Atan(float64(y2-y1)/float64(x2-x1))),
		starty:  StartingYPos(x1, y1, x2, y2, math.Atan(float64(y2-y1)/float64(x2-x1))),
		color: color.RGBA{
			R: 0xff,
			G: 0xff,
			B: 0xff,
			A: 0xff,
		},
	}
}

func Abs(x int) int {
	if x < 0 {
		return x * -1
	}
	return x
}

func DrawLine(img *ebiten.Image, x1, y1, x2, y2 int, c color.Color) {
	if Abs(y2-y1) <= Abs(x2-x1) {
		if x2 < x1 { // dx < 0
			x1, y1, x2, y2 = x2, y2, x1, y1
		}
		A := y2 - y1
		B := -(x2 - x1)
		// C := -B*y1 - A*x1  don't need C here
		dirY := 1
		if y2 < y1 { // dy < 0
			dirY = -1
			A = -A
		}
		d := A + B/2 // d0 which is the first midpoint (f(x1+1,y1+1/2))
		for x, y := x1, y1; x <= x2; x++ {
			img.Set(x, y, c)
			if d > 0 {
				y += dirY
				d += A + B
			} else {
				d += A
			}
		}
	} else {
		if y2 < y1 { // dy < 0
			x1, y1, x2, y2 = x2, y2, x1, y1
		}
		A := (x2 - x1)
		B := -(y2 - y1)
		// C := -B*x1 - A*y1  don't need C here
		dirX := 1
		if x2 < x1 { // dx < 0
			dirX = -1
			A = -A
		}
		d := A + B/2 // d0 which is the first midpoint (f(x1+1,y1+1/2))
		for x, y := x1, y1; y <= y2; y++ {
			img.Set(x, y, c)
			if d > 0 {
				x += dirX
				d += A + B
			} else {
				d += A
			}
		}
	}
}

func (l *Line) Update() {
	x := math.Cos(l.radians) * l.startx
	y := math.Sin(l.radians) * l.starty
	l.pos2.x = l.pos1.x + int(x)
	l.pos2.y = l.pos1.y + int(y)
	l.radians += math.Pi / 180
}

func (l *Line) Draw(screen *ebiten.Image) {
	DrawLine(screen, l.pos1.x, l.pos1.y, l.pos2.x, l.pos2.y, l.color)
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
		line:   NewLine(width/2, height/2, 130, 130),
	}
}

func (g *Game) Layout(outWidth, outHeight int) (w, h int) {
	return g.width, g.height
}

// Update updates a game state.
func (g *Game) Update() error {
	g.line.Update()
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

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
	pos1, pos2 Point
	radians    float64
	magnitude  int
	dir        int
	color      color.RGBA
}

func Abs(x int) int {
	if x < 0 {
		return x * -1
	}
	return x
}

func Magnitude(x1, y1, x2, y2 int) int {
	if Abs(x2-x1) > Abs(y2-y1) {
		return Abs(x2 - x1)
	}
	return Abs(y2 - y1)
}

func Dir(x1, y1, x2, y2 int) int {
	if x2 >= x1 {
		return 1
	}
	return -1
}

// NewLine initializes and returns a new Line instance.
func NewLine(x1, y1, x2, y2 int) *Line {
	return &Line{
		pos1:      Point{x: x1, y: y1},
		pos2:      Point{x: x2, y: y2},
		radians:   math.Atan(float64(y2-y1) / float64(x2-x1)), // math.Atan(k)  k выражено из уравнения прямой, проходящей через две точки
		dir:       Dir(x1, y1, x2, y2),
		magnitude: Magnitude(x1, y1, x2, y2),
		color: color.RGBA{
			R: 0xff,
			G: 0xff,
			B: 0xff,
			A: 0xff,
		},
	}
}

func DrawLine(img *ebiten.Image, x1, y1, x2, y2 int, c color.Color) {
	if Abs(y2-y1) <= Abs(x2-x1) {
		if x2 < x1 { // dx < 0
			x1, y1, x2, y2 = x2, y2, x1, y1
		}
		A := y2 - y1
		B := -(x2 - x1)
		C := -B*y1 - A*x1
		dirY := 1
		if y2 < y1 { // dy < 0
			dirY = -1
		}
		for x, y := x1, y1; x <= x2; x++ {
			img.Set(x, y, c)
			f := 2*A*(x+1) + B*(2*y+1*dirY) + 2*C // multiply everything by 2 to get rid of 0.5 because we only care about the sign
			if dirY*f > 0 {
				y += dirY
			}
		}
	} else {
		if y2 < y1 { // dy < 0
			x1, y1, x2, y2 = x2, y2, x1, y1
		}
		A := (x2 - x1)
		B := -(y2 - y1)
		C := -B*x1 - A*y1
		dirX := 1
		if x2 < x1 { // dx < 0
			dirX = -1
		}
		for x, y := x1, y1; y <= y2; y++ {
			img.Set(x, y, c)
			f := 2*A*(y+1) + B*(2*x+1*dirX) + 2*C // multiply everything by 2 to get rid of 0.5 because we only care about the sign
			if dirX*f > 0 {
				x += dirX
			}
		}
	}
}

func (l *Line) Update() {
	x := math.Cos(l.radians) * float64(l.magnitude)
	y := math.Sin(l.radians) * float64(l.magnitude)
	l.pos2.x = l.pos1.x + int(x)*l.dir
	l.pos2.y = l.pos1.y + int(y)*l.dir
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

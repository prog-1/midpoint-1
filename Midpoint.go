package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

//---------------------------Declaration--------------------------------

const (
	screenWidth  = 640
	screenHeight = 480
)

type Game struct {
	width, height int
	//here all the global variables are stored
}

//---------------------------Update-------------------------------------

func (g *Game) Update() error {
	//all logic on update
	//(can be divided on seperate functions for ex: "UpdateCircle")
	return nil
}

//---------------------------Draw-------------------------------------

func (g *Game) Draw(screen *ebiten.Image) {
	g.DrawLine(screen, screenWidth/2, screenHeight/2, screenWidth/2+100, screenHeight/2-100, color.RGBA{255, 255, 255, 255})
}

//-------------------------Functions----------------------------------

func (g *Game) DrawLine(screen *ebiten.Image, x1, y1, x2, y2 int, c color.Color) {

	//For debug
	screen.Set(x1, y1, c) //starting point
	screen.Set(x2, y2, c) //ending point

	s := 1 //sign of y (+ or -)
	if y2 < y1 {
		s = -1
	}

	A := y2 - y1      //Δy
	B := x1 - x2      // -Δx
	C := -B*y1 - A*x1 // C = Δx * y1 - Δy *x1

	fl := func(v int) float64 { return float64(v) } //to make formula shorter

	for x, y := x1, y1; x < x2; x++ {

		f := fl(A)*fl(x) + fl(B)*(fl(y)+(0.5*fl(s))) + fl(C) //Ax + By + C
		//B*y-0.5 to up, B*y+0.5 to down

		//if f < 0 fill the pixel on (x+1; y)
		if f >= 0 { //fill the pixel on (x+1; y+1)
			y = y + s //y-- to up, y++ to down
		}
		screen.Set(x, y, c) //filling the pixel

	}
}

//---------------------------Main-------------------------------------

func (g *Game) Layout(inWidth, inHeight int) (outWidth, outHeight int) {
	return g.width, g.height
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Midpoint")
	ebiten.SetWindowResizable(true)           //enablening window resize
	g := NewGame(screenWidth, screenHeight)   //creating game instance
	if err := ebiten.RunGame(g); err != nil { //running game
		log.Fatal(err)
	}
}

//New game instance creation
func NewGame(width, height int) *Game {
	return &Game{width: width, height: height}
}

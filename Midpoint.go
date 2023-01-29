package main

import (
	"image/color"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

//---------------------------Declaration--------------------------------

const (
	screenWidth  = 640
	screenHeight = 480
)

type Game struct {
	//here all the global variables are stored
	width, height int   //screen size
	l             *line //line struct
}

type line struct {
	x1, y1    int     //starting point
	x2, y2    int     //ending point
	magnitude float64 //line length
	radians   float64 //ending point angle
}

//---------------------------Update-------------------------------------

func (g *Game) Update() error {
	//all logic on update

	//increasing rotation
	g.l.radians += math.Pi / 90

	//to restrict the rotation
	// if g.l.radians > 2*math.Pi {
	// 	g.l.radians = 3 * math.Pi / 2
	// } else {
	// 	g.l.radians += math.Pi / 90
	// }

	return nil
}

//---------------------------Draw-------------------------------------

func (g *Game) Draw(screen *ebiten.Image) {

	//Ending point calculations with some fancy formulas
	x := int(math.Cos(g.l.radians) * g.l.magnitude)
	y := int(math.Sin(g.l.radians) * g.l.magnitude)

	//adding starting point to values
	g.l.x2 = g.l.x1 + x
	g.l.y2 = g.l.y1 + y

	//Line Draw
	g.DrawLine(screen, fl(g.l.x1), fl(g.l.y1), fl(g.l.x2), fl(g.l.y2), color.RGBA{255, 255, 255, 255}) //x2 = screenWidth/2+100 | y2 = screenHeight/2-100
}

//-------------------------Functions----------------------------------

func fl(v int) float64 { return float64(v) } //to make convertion shorter

func (g *Game) DrawLine(screen *ebiten.Image, x1, y1, x2, y2 float64, c color.Color) {

	if math.Abs(y2-y1) < math.Abs(x2-x1) {

		//drawing line from end to start if line is going left
		if x2 < x1 {
			x1, x2 = x2, x1
			y1, y2 = y2, y1
		}

		//For debug
		//screen.Set(int(x1), int(y1), c) //starting point
		//screen.Set(int(x2), int(y2), c) //ending point

		//Formula's variables
		A := y2 - y1 //Δy
		B := x1 - x2 // -Δx

		//Sign of y (+ or -)
		var s float64
		s = 1
		if y2 < y1 {
			s = -1
			A = -A
		}

		d := A + B/2 //d = d0// f from first middle point

		for x, y := x1, y1; x < x2; x++ {

			screen.Set(int(x), int(y), c) //filling the pixel

			//if f from d< 0 fill the pixel on (x+1; y)
			if d*s >= 0 { //fill the pixel on (x+1; y+1)
				y = y + s        //y-- to up, y++ to down
				d += (A + B) * s //dp+1 = dp+Δd  Δd = A+B
			} else {
				d += A * s //dp+1 = dp+Δd  Δd = A
			}

		}
	} else { //swapping x and y

		//drawing line from end to start if line is going left
		if y2 < y1 {
			x1, x2 = x2, x1
			y1, y2 = y2, y1
		}

		//For debug
		//screen.Set(int(x1), int(y1), c) //starting point
		//screen.Set(int(x2), int(y2), c) //ending point

		//Formula's variables
		A := x2 - x1 //Δy
		B := y1 - y2 // -Δx

		//Sign of x (+ or -)
		var s float64
		s = 1
		if x2 < x1 {
			s = -1
			A = -A
		}

		d := A + B/2 //d = d0// f from first middle point

		for x, y := x1, y1; y < y2; y++ {

			screen.Set(int(x), int(y), c) //filling the pixel

			//if f < 0 fill the pixel on (x+1; y)
			if d*s >= 0 { //fill the pixel on (x+1; y+1)
				x = x + s        //x-- to up, x++ to down
				d += (A + B) * s //dp+1 = dp+Δd  Δd = A+B
			} else {
				d += A * s //dp+1 = dp+Δd  Δd = A
			}

		}
	}

}

//---------------------------Main-------------------------------------

func (g *Game) Layout(inWidth, inHeight int) (outWidth, outHeight int) {
	return g.width, g.height
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Midpoint")
	ebiten.SetWindowResizable(true) //enablening window resize

	//creating game instance
	g := &Game{width: screenWidth, height: screenHeight,
		l: &line{x1: screenWidth / 2, y1: screenHeight / 2, x2: 100, y2: 0, magnitude: 150, radians: 3 * math.Pi / 2 /*0*/}} //configuring line

	//running game
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}

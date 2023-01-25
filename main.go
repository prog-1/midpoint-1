package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	Wwidth  = 700
	Wheight = 700
)

type Game struct {
	width, height int
}

var c = color.RGBA{244, 212, 124, 255}

func (g *Game) DrawLine(img *ebiten.Image, x1, y1, x2, y2 float64, c color.Color) {
	if x2 < x1 {
		x1, x2 = x2, x1
		y1, y2 = y2, y1
	}
	dx, dy := x2-x1, y2-y1
	xp, yp := x1, y1
	if dx > dy {
		A, B, C := dy, -dx, dx*y1-dy*x1
		dirY := 1.0
		if dy < 0 {
			dirY = -1
		}
		for xp != x2 && yp != y2 {
			// if dirY*(A*(xp+1)+B*(yp+0.5)+C) <= 0 {
			// 	xp, yp = xp+1, yp+0.5
			// 	img.Set(int(xp), int(yp), c)
			// }
			if dirY*(A*(xp+1)+B*(yp+0.5)+C) > 0 {
				xp, yp = xp+1, yp+dirY
				img.Set(int(xp), int(yp), c)
			}
		}
	}
	if dx < dy {
		A, B, C := dx, -dy, dy*x1-dx*y1
		dirX := 1.0
		if dx < 0 {
			dirX = -1
		}
		for xp != x2 && yp != y2 {
			// if dirX*(A*(xp+0.5)+B*(yp+1)+C) <= 0 {
			// 	xp, yp = xp+0.5, yp+1
			// 	img.Set(int(xp), int(yp), c)
			// }
			if dirX*(A*(xp+0.5)+B*(yp+1)+C) > 0 {
				xp, yp = xp+dirX, yp+1
				img.Set(int(xp), int(yp), c)
			}
		}
	}
}

func (g *Game) Draw(screen *ebiten.Image) {
	//g.DrawLine(screen, 100, 600, 200, 300, c)
	g.DrawLine(screen, 100, 300, 200, 500, c)

}

func (g *Game) Update() error {
	return nil

}
func (g *Game) Layout(int, int) (w, h int) {
	return g.width, g.height
}

func main() {
	ebiten.SetWindowSize(Wwidth, Wheight)
	if err := ebiten.RunGame(NewGame(Wwidth, Wheight)); err != nil {
		log.Fatal(err)
	}
}
func NewGame(width, height int) *Game {
	return &Game{width: width, height: height}
}

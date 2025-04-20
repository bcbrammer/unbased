package main

import (
	"fmt"
	"image"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	isDebugMode bool = true
)

type Anim struct {
	CurrentImg   *ebiten.Image
	CurrentIndex int
	Imgs         []*ebiten.Image
}

func (a *Anim) setCurrentImg() error {
	if a.CurrentIndex+1 == len(a.Imgs) {
		a.CurrentIndex = 0
	} else {
		a.CurrentIndex++
	}
	a.CurrentImg = a.Imgs[a.CurrentIndex]
	return nil
}

type Sprite struct {
	Anim *Anim
	X    float64
	Y    float64
}

type Game struct {
	Player *Entity
	Ents   []*Entity
}

type Entity struct {
	*Sprite
}

// defaults to 60 ticks per sec
func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		return ebiten.Termination
	}

	movePlayer(g)

	return nil
}

func movePlayer(g *Game) {
	// W
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		if ebiten.IsKeyPressed(ebiten.KeyA) {
			g.Player.X -= 1
			g.Player.Y -= 1
			return
		} else if ebiten.IsKeyPressed(ebiten.KeyD) {
			g.Player.X += 1
			g.Player.Y -= 1
			return
		} else {
			g.Player.Y -= 2
		}
	}
	// S
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		if ebiten.IsKeyPressed(ebiten.KeyA) {
			g.Player.X -= 1
			g.Player.Y += 1
		} else if ebiten.IsKeyPressed(ebiten.KeyD) {
			g.Player.X += 1
			g.Player.Y += 1
		} else {
			g.Player.Y += 2
		}
	}
	// A
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		if ebiten.IsKeyPressed(ebiten.KeyW) {
			g.Player.X -= 1
			g.Player.Y -= 1
		} else if ebiten.IsKeyPressed(ebiten.KeyS) {
			g.Player.X -= 1
			g.Player.Y += 1
		} else {
			g.Player.X -= 2
		}
	}
	// D
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		if ebiten.IsKeyPressed(ebiten.KeyW) {
			g.Player.X += 1
			g.Player.Y -= 1
		} else if ebiten.IsKeyPressed(ebiten.KeyS) {
			g.Player.X += 1
			g.Player.Y += 1
		} else {
			g.Player.X += 2
		}
	}

	// Keep player in bounds
	if g.Player.X < 0 {
		g.Player.X = 0
	}
	if g.Player.X > float64(240) {
		g.Player.X = float64(240)
	}
	if g.Player.Y < 0 {
		g.Player.Y = 0
	}
	if g.Player.Y > float64(320) {
		g.Player.Y = float64(320)
	}
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{120, 180, 255, 255})
	if isDebugMode {
		ebitenutil.DebugPrint(screen, "Debug Mode")
	}

	pdio := ebiten.DrawImageOptions{}
	pdio.GeoM.Translate(g.Player.X, g.Player.Y)

	screen.DrawImage(
		g.Player.Anim.CurrentImg.SubImage(image.Rect(0, 0, 128, 128)).(*ebiten.Image),
		&pdio,
	)

	dios := ebiten.DrawImageOptions{}
	for _, ent := range g.Ents {
		dios.GeoM.Translate(ent.X, ent.Y)

		screen.DrawImage(
			ent.Anim.CurrentImg.SubImage(image.Rect(0, 0, 128, 128)).(*ebiten.Image),
			&dios,
		)
		dios.GeoM.Reset()
	}
}

func (g *Game) Layout(outsideWidth int, outsideHeight int) (screenwidth int, screenHeight int) {
	return 320, 240
}

func main() {
	fmt.Println("main() init")
	ebiten.SetWindowSize(1280, 720)
	ebiten.SetWindowTitle("td0.0.3")

	playerImg, _, err := ebitenutil.NewImageFromFile("assets/images/player/idle0.png")
	if err != nil {
		log.Fatalf("game err: %v", err)
	}

	g := &Game{
		Player: &Entity{
			&Sprite{
				Anim: &Anim{
					CurrentImg:   playerImg,
					CurrentIndex: 0,
					Imgs:         []*ebiten.Image{playerImg},
				},
				X: 100,
				Y: 100,
			},
		},
		Ents: []*Entity{
			{
				&Sprite{
					Anim: &Anim{
						CurrentImg:   playerImg,
						CurrentIndex: 0,
						Imgs:         []*ebiten.Image{playerImg},
					},
					X: 100,
					Y: 100,
				},
			},
			{
				&Sprite{
					Anim: &Anim{
						CurrentImg:   playerImg,
						CurrentIndex: 0,
						Imgs:         []*ebiten.Image{playerImg},
					},
					X: 50,
					Y: 50,
				},
			},
		},
	}

	if err := ebiten.RunGame(g); err != nil {
		log.Fatalf("game err: %v", err)
	}
}

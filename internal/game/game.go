package game

import (
	"fmt"
	"image/color"

	"github.com/bcbrammer/unbased/internal/engine"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

// Game represents our game state and implements the Engine.Game interface
type Game struct {
	Engine     *engine.Engine
	PlayerX    float64
	PlayerY    float64
	PlayerSize float64
	Speed      float64
}

// NewGame creates a new game instance
func NewGame(e *engine.Engine) *Game {
	return &Game{
		Engine:     e,
		PlayerX:    float64(e.Width) / 2,
		PlayerY:    float64(e.Height) / 2,
		PlayerSize: 40,
		Speed:      4.0,
	}
}

// Run starts the game
func (g *Game) Run() error {
	return g.Engine.Run(g)
}

// Update handles game logic and input
func (g *Game) Update() error {
	// Exit on ESC
	if g.Engine.IsKeyJustPressed(ebiten.KeyEscape) {
		return ebiten.Termination
	}

	// Movement
	if g.Engine.IsKeyPressed(ebiten.KeyW) || g.Engine.IsKeyPressed(ebiten.KeyUp) {
		g.PlayerY -= g.Speed
	}
	if g.Engine.IsKeyPressed(ebiten.KeyS) || g.Engine.IsKeyPressed(ebiten.KeyDown) {
		g.PlayerY += g.Speed
	}
	if g.Engine.IsKeyPressed(ebiten.KeyA) || g.Engine.IsKeyPressed(ebiten.KeyLeft) {
		g.PlayerX -= g.Speed
	}
	if g.Engine.IsKeyPressed(ebiten.KeyD) || g.Engine.IsKeyPressed(ebiten.KeyRight) {
		g.PlayerX += g.Speed
	}

	// Keep player in bounds
	if g.PlayerX < 0 {
		g.PlayerX = 0
	}
	if g.PlayerX > float64(g.Engine.Width) {
		g.PlayerX = float64(g.Engine.Width)
	}
	if g.PlayerY < 0 {
		g.PlayerY = 0
	}
	if g.PlayerY > float64(g.Engine.Height) {
		g.PlayerY = float64(g.Engine.Height)
	}

	return nil
}

// Draw renders the game
func (g *Game) Draw(screen *ebiten.Image) {
	// Clear screen with dark background
	screen.Fill(color.RGBA{20, 20, 30, 255})

	// Draw player as red rectangle with white outline
	vector.DrawFilledRect(screen,
		float32(g.PlayerX-g.PlayerSize/2),
		float32(g.PlayerY-g.PlayerSize/2),
		float32(g.PlayerSize),
		float32(g.PlayerSize),
		color.RGBA{255, 0, 0, 255}, // Red fill
		true,
	)

	// Draw outline
	vector.StrokeRect(screen,
		float32(g.PlayerX-g.PlayerSize/2),
		float32(g.PlayerY-g.PlayerSize/2),
		float32(g.PlayerSize),
		float32(g.PlayerSize),
		2,                              // Line width
		color.RGBA{255, 255, 255, 255}, // White outline
		true,
	)

	// TODO: debuggin
	ebitenutil.DebugPrint(
		screen,
		fmt.Sprintf("FPS: %0.1f\nPos: (%0.1f, %0.1f)",
			ebiten.ActualFPS(),
			g.PlayerX,
			g.PlayerY,
		),
	)
}

// Layout defines the game's logical screen size
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.Engine.Width, g.Engine.Height
}

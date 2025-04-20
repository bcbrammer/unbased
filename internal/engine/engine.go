package engine

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Engine struct {
	Width   int
	Height  int
	Title   string
	Running bool
}

func NewEngine(width int, height int, title string) (*Engine, error) {
	return &Engine{
		Width:   width,
		Height:  height,
		Title:   title,
		Running: true,
	}, nil
}

func (e *Engine) IsRunning() bool {
	return e.Running
}

func (e *Engine) SetRunning(Running bool) {
	e.Running = Running
}

func (e *Engine) IsKeyPressed(key ebiten.Key) bool {
	return ebiten.IsKeyPressed(key)
}

func (e *Engine) IsKeyJustPressed(key ebiten.Key) bool {
	return inpututil.IsKeyJustPressed(key)
}

func (e *Engine) Run(game Game) error {
	ebiten.SetWindowSize(e.Width, e.Height)
	ebiten.SetWindowTitle(e.Title)
	return ebiten.RunGame(game)
}

// Game is the interface that must be implemented by your game
type Game interface {
	Update() error
	Draw(screen *ebiten.Image)
	// Layout takes the outside size (e.g., the window size) and returns the logical game size
	Layout(outsideWidth int, outsideHeight int) (int, int)
}

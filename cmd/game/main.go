package main

import (
	"fmt"
	"log"

	"github.com/bcbrammer/unbased/internal/engine"
	"github.com/bcbrammer/unbased/internal/game"
)

func main() {
	fmt.Println("main() init")

	eng, err := engine.NewEngine(800, 600, "td0.0.2")
	if err != nil {
		log.Fatalf("failed to construct engine: %v", err)
	}

	g := game.NewGame(eng)
	if err := g.Run(); err != nil {
		log.Fatalf("game err: %v", err)
	}
}

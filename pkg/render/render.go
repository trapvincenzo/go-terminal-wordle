package render

import "github.com/trapvincenzo/go-terminal-wordle/pkg/game"

type Renderer interface {
	Render(*game.Game)
	Print(string)
}

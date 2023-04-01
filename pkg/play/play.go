package play

import (
	"github.com/trapvincenzo/go-terminal-wordle/pkg/game"
	"github.com/trapvincenzo/go-terminal-wordle/pkg/reader"
	"github.com/trapvincenzo/go-terminal-wordle/pkg/render"
)

func Start(g *game.Game, render render.Renderer, reader reader.Reader) {

	for {
		if g.State != game.InProgress {
			break
		}

		render.Render(g)

		if g.State == game.InProgress {
			text := ""
			for {
				render.Print("Try a 5 letters word: ")
				text = reader.ReadString()

				if len(text) == 5 {
					break
				}
			}

			g.TryNewWord(text)
		}
	}

	render.Render(g) // render final status
}

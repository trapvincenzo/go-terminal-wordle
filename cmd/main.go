package main

import (
	"os"

	"github.com/trapvincenzo/go-terminal-wordle/pkg/game"
	"github.com/trapvincenzo/go-terminal-wordle/pkg/play"
	"github.com/trapvincenzo/go-terminal-wordle/pkg/reader"
	"github.com/trapvincenzo/go-terminal-wordle/pkg/render"
	"github.com/trapvincenzo/go-terminal-wordle/pkg/storage"
)

func main() {
	g := game.NewGame(&storage.StaticStorage{})
	play.Start(g, &render.PrintRender{Writer: os.Stdout}, &reader.Bufio{})
}

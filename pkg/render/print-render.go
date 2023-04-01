package render

import (
	"fmt"
	"io"
	"sort"

	"github.com/fatih/color"
	"github.com/trapvincenzo/go-terminal-wordle/pkg/game"
)

type PrintRender struct {
	Writer io.Writer
}

var colorMap = map[game.BoxStatus]color.Attribute{
	game.NotFound:     color.Reset,
	game.CorrectPlace: color.BgGreen,
	game.Misplaced:    color.BgYellow,
}

func (p PrintRender) Render(g *game.Game) {
	clearScreen()

	renderTitle(g, &p)
	renderBoard(g, &p)
	renderLetters(g, &p)
	readerState(g, &p)
}

func (p PrintRender) Print(s string) {
	fmt.Print(s)
}

func renderBoard(g *game.Game, p *PrintRender) {
	for i := 0; i < game.Rows; i++ {
		fmt.Fprintln(p.Writer, "+---+---+---+---+---+")
		for j := 0; j < game.Cols; j++ {

			fmt.Fprint(p.Writer, "|")
			c := color.New(color.Attribute(colorMap[g.GameMatrix[i][j].Status]))
			c.Fprintf(p.Writer, " %s ", g.GameMatrix[i][j].Value)
			if j == game.Cols-1 {
				fmt.Fprintln(p.Writer, "|")
			}
		}
	}
	fmt.Fprintln(p.Writer, "+---+---+---+---+---+")
	fmt.Fprint(p.Writer, "\n")
}

func readerState(g *game.Game, p *PrintRender) {
	if g.State == game.Won {
		fmt.Fprintln(p.Writer, "   +=============+")
		fmt.Fprintln(p.Writer, "   |   YOU WON   |")
		fmt.Fprintln(p.Writer, "   +=============+")
		fmt.Fprint(p.Writer, "\n")
	}

	if g.State == game.Lost {
		fmt.Fprintln(p.Writer, "   +------------+")
		fmt.Fprintln(p.Writer, "   |  YOU LOST  |")
		fmt.Fprintln(p.Writer, "   +------------+")
		fmt.Fprint(p.Writer, "\n")
	}
}

func renderTitle(g *game.Game, p *PrintRender) {
	fmt.Fprintln(p.Writer, "   +=============+")
	fmt.Fprintf(p.Writer, "   | %s |\n", g.Title)
	fmt.Fprintln(p.Writer, "   +=============+")
	fmt.Fprint(p.Writer, "\n")
}

func renderLetters(g *game.Game, p *PrintRender) {
	if g.State != game.InProgress {
		return
	}

	fmt.Fprint(p.Writer, "Letters: ")

	f := color.New(color.Faint).Add(color.CrossedOut)

	for k, v := range getSortedLettersMap(g) {
		if k%8 == 0 {
			fmt.Fprintln(p.Writer, "")
		}

		l := g.AvailableLetters[v]

		if l == game.Off {
			f.Fprintf(p.Writer, " %s ", v)
			continue
		}
		fmt.Fprintf(p.Writer, " %s ", v)

	}

	fmt.Fprintln(p.Writer, "")
	fmt.Fprintln(p.Writer, "")

}

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

func getSortedLettersMap(g *game.Game) []string {

	keys := make([]string, 0, len(g.AvailableLetters))
	for k := range g.AvailableLetters {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

package game

import (
	"strings"

	"github.com/trapvincenzo/go-terminal-wordle/pkg/storage"
)

const Rows = 5
const Cols = 5

type Game struct {
	AvailableLetters map[string]LetterStatus
	GameMatrix       [Rows][Cols]Box
	State            State
	Title            string

	cursor             int
	currentWord        string
	currentWordIndexed []string
	currentLetterCount map[string]int
}

type Wordle interface {
	TryNewWord(word string)
}

type Box struct {
	Value  string
	Status BoxStatus
}

type Letter struct {
	Value  string
	Status LetterStatus
}

type BoxStatus int

const (
	CorrectPlace BoxStatus = iota
	Misplaced
	NotFound
)

type State int

const (
	InProgress State = iota
	Won
	Lost
)

type LetterStatus int

const (
	On LetterStatus = iota
	Off
)

func NewGame(storage storage.Storage) *Game {
	matrix := [Rows][Cols]Box{}
	letters := map[string]LetterStatus{}

	for i := 0; i < Rows; i++ {
		for j := 0; j < Cols; j++ {
			matrix[i][j] = Box{Value: " ", Status: NotFound}
		}
	}

	for i := 65; i <= 90; i++ {
		letters[string(rune(i))] = On
	}

	word := storage.GetWord()

	lettersCount := map[string]int{}

	wordIndexed := strings.Split(word, "")
	for _, v := range wordIndexed {
		lettersCount[v]++
	}

	return &Game{
		GameMatrix:         matrix,
		AvailableLetters:   letters,
		Title:              "TERM-WORDLE",
		currentWord:        word,
		currentWordIndexed: strings.Split(word, ""),
		currentLetterCount: lettersCount,
	}
}

func (g *Game) TryNewWord(word string) {
	g.addWordToTheMatrix(word)

	g.cursor++

	if g.cursor == Rows {
		// this is the last round,
		// did the player won?
		if g.State != Won {
			g.State = Lost
		}
	}
}

func (g *Game) addWordToTheMatrix(word string) {

	letterStatus, lettersFound, completed := g.checkWord(word)

	wordIndexed := strings.Split(word, "")

	for i := 0; i < Cols; i++ {
		status := letterStatus[i]
		if status == Misplaced && strings.Contains(lettersFound, wordIndexed[i]) {
			status = NotFound
		}

		g.GameMatrix[g.cursor][i].Value = wordIndexed[i]
		g.GameMatrix[g.cursor][i].Status = status
	}

	if completed {
		g.State = Won
	}
}

func (g *Game) checkWord(word string) (map[int]BoxStatus, string, bool) {
	letterCount := map[string]int{}
	letterStatus := map[int]BoxStatus{}
	lettersFound := ""
	count := 0

	wordIndexed := strings.Split(word, "")
	for k, v := range wordIndexed {
		if !strings.Contains(g.currentWord, v) {
			letterStatus[k] = NotFound
			g.AvailableLetters[strings.ToUpper(v)] = Off
			continue
		}

		if v == g.currentWordIndexed[k] {
			letterStatus[k] = CorrectPlace
			letterCount[v] = g.currentLetterCount[v]
			lettersFound += v
			count++
			continue
		}

		if letterCount[v] < g.currentLetterCount[v] {
			letterCount[v]++
			letterStatus[k] = Misplaced
			continue
		}

		letterStatus[k] = NotFound
	}

	return letterStatus, lettersFound, count == 5
}

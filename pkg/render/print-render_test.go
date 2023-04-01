package render

import (
	"bytes"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/trapvincenzo/go-terminal-wordle/pkg/game"
	"github.com/trapvincenzo/go-terminal-wordle/pkg/storage"
)

type UnitSuite struct {
	suite.Suite
	newGame *game.Game
}

func TestUnitSuite(t *testing.T) {
	suite.Run(t, &UnitSuite{})
}

func (u *UnitSuite) SetupTest() {
	s := storage.NewMockStorage(u.T())
	s.EXPECT().GetWord().Return("maybe")
	u.newGame = game.NewGame(s)
}

func (u *UnitSuite) TestTitleIsRenderedCorrectly() {
	var output bytes.Buffer
	render := &PrintRender{&output}

	renderTitle(u.newGame, render)

	assert.Equal(u.T(), readFixture("title"), output.String())
}

func (u *UnitSuite) TestEmptyBoardIsRenderedCorrectly() {
	var output bytes.Buffer

	render := &PrintRender{&output}

	renderBoard(u.newGame, render)

	assert.Equal(u.T(), readFixture("emptyBoard"), output.String())
}

func (u *UnitSuite) TestBoardIsRenderedCorrectly() {
	var output bytes.Buffer
	render := &PrintRender{&output}

	u.newGame.TryNewWord("tests")
	u.newGame.TryNewWord("maybe")
	renderBoard(u.newGame, render)

	assert.Equal(u.T(), readFixture("board"), output.String())
}

func (u *UnitSuite) TestStateRenderedCorrectly() {
	var output bytes.Buffer
	render := &PrintRender{&output}

	u.newGame.State = game.Lost
	readerState(u.newGame, render)

	assert.Equal(u.T(), readFixture("lost"), output.String())

	var output2 bytes.Buffer
	render = &PrintRender{&output2}

	u.newGame.State = game.Won
	readerState(u.newGame, render)

	assert.Equal(u.T(), readFixture("won"), output2.String())
}

func (u *UnitSuite) TestLettersRenderedCorrectly() {
	var output bytes.Buffer
	render := &PrintRender{&output}

	renderLetters(u.newGame, render)

	assert.Equal(u.T(), readFixture("letters"), output.String())
}

func readFixture(name string) string {
	dat, err := os.ReadFile(fmt.Sprintf("./fixtures/%s.txt", name))
	if err != nil {
		panic(err)
	}
	return string(dat)
}

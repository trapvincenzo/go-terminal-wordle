package play

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"github.com/trapvincenzo/go-terminal-wordle/pkg/game"
	"github.com/trapvincenzo/go-terminal-wordle/pkg/reader"
	"github.com/trapvincenzo/go-terminal-wordle/pkg/render"
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
	newGame := game.NewGame(s)

	u.newGame = newGame
}

func (u *UnitSuite) TestGameIsOverIfTheGameIsNotInProgress() {

	mockRender := &render.MockRenderer{}
	mockReader := &reader.MockReader{}

	u.newGame.State = game.Won
	mockRender.On("Render", mock.Anything).Return("")
	Start(u.newGame, mockRender, mockReader)
	mockRender.AssertNotCalled(u.T(), "Print", "Try a 5 letters word: ")
}

func (u *UnitSuite) TestWordIsReadProperlyAndPassedToTheGame() {
	mockRender := &render.MockRenderer{}
	mockReader := &reader.MockReader{}

	mockRender.On("Render", mock.Anything).Return("")
	mockRender.On("Print", "Try a 5 letters word: ").Return("")
	mockReader.On("ReadString").Return("tests")

	Start(u.newGame, mockRender, mockReader)

	word := [5]game.Box{
		{Value: "t", Status: game.NotFound},
		{Value: "e", Status: game.Misplaced},
		{Value: "s", Status: game.NotFound},
		{Value: "t", Status: game.NotFound},
		{Value: "s", Status: game.NotFound},
	}
	assert.Equal(u.T(), word, u.newGame.GameMatrix[0])

}

package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/trapvincenzo/go-terminal-wordle/pkg/storage"
)

type UnitSuite struct {
	suite.Suite
	newGame *Game
}

func (u *UnitSuite) SetupTest() {
	s := storage.NewMockStorage(u.T())
	s.EXPECT().GetWord().Return("maybe")
	newGame := NewGame(s)

	u.newGame = newGame
}

func (u *UnitSuite) TestNewGameIsInitialsedCorrectly() {

	assert.Equal(u.T(), u.newGame.currentWord, "maybe")
	assert.Equal(u.T(), u.newGame.currentWordIndexed, []string{"m", "a", "y", "b", "e"})
	assert.Equal(u.T(), 26, len(u.newGame.AvailableLetters))
}

func (u *UnitSuite) TestCursorMovesDown() {
	assert.Equal(u.T(), 0, u.newGame.cursor)
	u.newGame.TryNewWord("tests")
	assert.Equal(u.T(), 1, u.newGame.cursor)
	u.newGame.TryNewWord("tests")
	assert.Equal(u.T(), 2, u.newGame.cursor)
}

func (u *UnitSuite) TestWhenCursorReachesMaxRowGameEnds() {
	u.newGame.TryNewWord("tests")
	u.newGame.TryNewWord("tests")
	u.newGame.TryNewWord("tests")
	u.newGame.TryNewWord("tests")
	u.newGame.TryNewWord("tests")

	assert.Equal(u.T(), 5, u.newGame.cursor)
	assert.Equal(u.T(), Lost, u.newGame.State)
}

func (u *UnitSuite) TestWhenTheWordIsCorrectTheGameEndsAsWinner() {
	assert.Equal(u.T(), InProgress, u.newGame.State)
	u.newGame.TryNewWord("maybe")
	assert.Equal(u.T(), Won, u.newGame.State)
}

func (u *UnitSuite) TestWhenTheWordIsCorrectAtTheLastTryTheGameEndsAsWinner() {
	assert.Equal(u.T(), InProgress, u.newGame.State)
	u.newGame.TryNewWord("mayba")
	u.newGame.TryNewWord("maybr")
	u.newGame.TryNewWord("maybt")
	u.newGame.TryNewWord("mayby")
	u.newGame.TryNewWord("maybe")
	assert.Equal(u.T(), Won, u.newGame.State)
}

func (u *UnitSuite) TestCheckWordTestCases() {
	cases := tableCheckWordTestCases()
	for _, v := range cases {
		boxStatus, lettersFound, completed := u.newGame.checkWord(v.word)
		assert.Equal(u.T(), v.expected, boxStatus)
		assert.Equal(u.T(), v.lettersFound, lettersFound)
		assert.Equal(u.T(), v.completed, completed)
	}
}

func tableCheckWordTestCases() []struct {
	word         string
	expected     map[int]BoxStatus
	completed    bool
	lettersFound string
} {
	return []struct {
		word         string
		expected     map[int]BoxStatus
		completed    bool
		lettersFound string
	}{
		{
			completed: false,
			word:      "tests",
			expected: map[int]BoxStatus{
				0: NotFound,
				1: Misplaced,
				2: NotFound,
				3: NotFound,
				4: NotFound,
			},
			lettersFound: "",
		},
		{
			completed: true,
			word:      "maybe",
			expected: map[int]BoxStatus{
				0: CorrectPlace,
				1: CorrectPlace,
				2: CorrectPlace,
				3: CorrectPlace,
				4: CorrectPlace,
			},
			lettersFound: "maybe",
		},
		{
			// too 'a's only one should be flagged
			completed: false,
			word:      "mayba",
			expected: map[int]BoxStatus{
				0: CorrectPlace,
				1: CorrectPlace,
				2: CorrectPlace,
				3: CorrectPlace,
				4: NotFound,
			},
			lettersFound: "mayb",
		},
		{
			// Only 1 a and m should be flagged
			// as misplaces, the othe as not found
			word: "amama",
			expected: map[int]BoxStatus{
				0: Misplaced,
				1: Misplaced,
				2: NotFound,
				3: NotFound,
				4: NotFound,
			},
			lettersFound: "",
		},
	}
}

func (u *UnitSuite) TestWordIsAddedToTheMatrixWithTheRightStatus() {
	u.newGame.addWordToTheMatrix("tests")

	entry := [5]Box{
		{Value: "t", Status: NotFound},
		{Value: "e", Status: Misplaced},
		{Value: "s", Status: NotFound},
		{Value: "t", Status: NotFound},
		{Value: "s", Status: NotFound},
	}

	assert.Equal(u.T(), entry, u.newGame.GameMatrix[0])
}

func (u *UnitSuite) TestLettersNotFoundAreOff() {
	u.newGame.addWordToTheMatrix("tests")

	assert.Equal(u.T(), Off, u.newGame.AvailableLetters["T"])
	assert.Equal(u.T(), Off, u.newGame.AvailableLetters["S"])
}

func (u *UnitSuite) TestLettersFoundAreStillOn() {
	u.newGame.addWordToTheMatrix("tests")

	assert.Equal(u.T(), On, u.newGame.AvailableLetters["E"])
}

func TestUnitSuite(t *testing.T) {
	suite.Run(t, &UnitSuite{})
}

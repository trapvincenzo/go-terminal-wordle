package reader

import (
	"bufio"
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type UnitSuite struct {
	suite.Suite
}

func TestUnitSuite(t *testing.T) {
	suite.Run(t, &UnitSuite{})
}

func (u *UnitSuite) TestInputAreCorrectlyRead() {
	var input bytes.Buffer
	reader := &Bufio{
		reader: bufio.NewReader(&input),
	}
	input.Write([]byte("maybe"))
	assert.Equal(u.T(), "maybe", reader.ReadString())

}

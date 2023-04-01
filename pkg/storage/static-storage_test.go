package storage

import (
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

func (u *UnitSuite) TestGoWordReturnTheCorrectWord() {
	s := &StaticStorage{}
	assert.Equal(u.T(), "maybe", s.GetWord())
}

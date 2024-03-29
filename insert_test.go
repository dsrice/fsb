package fsb_test

import (
	"fsb"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type InsertSuite struct {
	suite.Suite
}

// Test_Insert is a unit test for the Insert method
func (s *InsertSuite) Test_Insert() {
	sb := fsb.Insert("id").Into(fsb.Table("users")).Value(1)
	sql, err := sb.ToSQL()

	assert.Equal(s.T(), "INSERT INTO users ( id ) VALUES ( 1 );", sql)
	assert.Nil(s.T(), err)
}

// Test_InsertMulti is a unit test for the InsertMulti method
func (s *InsertSuite) Test_InsertMulti() {
	sb := fsb.Insert("id", "name").Into(fsb.Table("users")).Value(1, "test")
	sql, err := sb.ToSQL()

	assert.Equal(s.T(), "INSERT INTO users ( id, name ) VALUES ( 1, 'test' );", sql)
	assert.Nil(s.T(), err)
}

// Test_InsertCol is a unit test for the InsertCol method
func (s *InsertSuite) Test_InsertCol() {
	user := fsb.Table("users")
	sb := fsb.Insert(user.Col("id")).Into(user).Value(1)
	sql, err := sb.ToSQL()

	assert.Equal(s.T(), "INSERT INTO users ( users.id ) VALUES ( 1 );", sql)
	assert.Nil(s.T(), err)
}

func (s *InsertSuite) Test_InsertMultiValue() {
	user := fsb.Table("users")
	sb := fsb.Insert(user.Col("id")).Into(user).Value(1).Value(2)
	sql, err := sb.ToSQL()

	assert.Equal(s.T(), "INSERT INTO users ( users.id ) VALUES ( 1 ) , ( 2 );", sql)
	assert.Nil(s.T(), err)
}

func TestInsertSuite(t *testing.T) {
	suite.Run(t, new(InsertSuite))
}

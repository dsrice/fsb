package fsb_test

import (
	"fsb"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type SelectSuite struct {
	suite.Suite
}

// Test_SelectString tests the SelectString method in the SelectSuite struct.
func (s *SelectSuite) Test_SelectString() {
	sb := fsb.Select().From(fsb.Table("users"))
	sql, err := sb.ToSQL()

	assert.Equal(s.T(), "SELECT * FROM users;", sql)
	assert.Nil(s.T(), err)
}

// Test_SelectString_Field tests the SelectString method in the SelectSuite struct.
// It selects the "id" field from the "users" table and generates a SQL query.
// The generated SQL query should be "SELECT id FROM users;".
// It asserts that the generated SQL query and error are as expected.
func (s *SelectSuite) Test_SelectString_Field() {
	sb := fsb.Select("id").From(fsb.Table("users"))
	sql, err := sb.ToSQL()

	assert.Equal(s.T(), "SELECT id FROM users;", sql)
	assert.Nil(s.T(), err)
}

// Test_SelectString_WhereInt tests the SelectString_WhereInt method in the SelectSuite struct.
func (s *SelectSuite) Test_SelectString_WhereInt() {
	sb := fsb.Select().
		From(fsb.Table("users")).
		Where(fsb.Eq("id", 1))

	sql, err := sb.ToSQL()

	assert.Equal(s.T(), "SELECT * FROM users WHERE id = 1;", sql)
	assert.Nil(s.T(), err)
}

// Test_SelectString_WhereString tests the SelectString_WhereString method in the SelectSuite struct.
func (s *SelectSuite) Test_SelectString_WhereString() {
	sb := fsb.Select().
		From(fsb.Table("users")).
		Where(fsb.Eq("name", "test"))

	sql, err := sb.ToSQL()

	assert.Equal(s.T(), "SELECT * FROM users WHERE name = 'test';", sql)
	assert.Nil(s.T(), err)
}

// Test_SelectString_WhereAND tests the WhereAND method in the SelectSuite struct.
// It checks if the generated SQL query from SelectString_WhereAnd method matches the expected SQL query.
func (s *SelectSuite) Test_SelectString_WhereAND() {
	sb := fsb.Select().
		From(fsb.Table("users")).
		Where(fsb.Eq("name", "test").AND(fsb.Eq("id", 1)))

	sql, err := sb.ToSQL()

	assert.Equal(s.T(), "SELECT * FROM users WHERE name = 'test' AND id = 1;", sql)
	assert.Nil(s.T(), err)
}

// Test_SelectString_WhereOR tests the WhereOR method in the SelectSuite struct.
func (s *SelectSuite) Test_SelectString_WhereOR() {
	sb := fsb.Select().
		From(fsb.Table("users")).
		Where(fsb.Eq("name", "test").OR(fsb.Eq("id", 1)))

	sql, err := sb.ToSQL()

	assert.Equal(s.T(), "SELECT * FROM users WHERE name = 'test' OR id = 1;", sql)
	assert.Nil(s.T(), err)
}

// Test_SelectString_WhereAndOr tests the SelectString method in the SelectSuite struct.
// It creates a SelectContainer object and calls the From method to set the table to "users".
// Then it calls the Where method to set the conditions to "name = 'test' AND (id = 1 OR id = 2)".
// Finally, it calls the ToSQL method to generate the SQL string and checks if it matches the expected value.
// The test also checks if there are no errors returned.
func (s *SelectSuite) Test_SelectString_WhereAndOr() {
	sb := fsb.Select().
		From(fsb.Table("users")).
		Where(fsb.Eq("name", "test").
			AND(
				fsb.Eq("id", 1).OR(fsb.Eq("id", 2)),
			),
		)

	sql, err := sb.ToSQL()

	assert.Equal(s.T(), "SELECT * FROM users WHERE name = 'test' AND (id = 1 OR id = 2);", sql)
	assert.Nil(s.T(), err)
}

// Test_SelectString_WhereOrAnd tests the WhereOrAnd method in the SelectSuite struct.
func (s *SelectSuite) Test_SelectString_WhereOrAnd() {
	sb := fsb.Select().
		From(fsb.Table("users")).
		Where(fsb.Eq("name", "test").
			OR(
				fsb.Eq("id", 1).AND(fsb.Eq("id", 2)),
			),
		)

	sql, err := sb.ToSQL()

	assert.Equal(s.T(), "SELECT * FROM users WHERE name = 'test' OR (id = 1 AND id = 2);", sql)
	assert.Nil(s.T(), err)
}

// Test_SelectString_WhereMulti tests the SelectString_WhereMulti method in the SelectSuite struct.
// Selects from the "users" table with the following conditions:
// - "name" equals "test" OR ("id" equals 1 AND "id" equals 2)
// - AND ("login_id" equals "test2" AND "name" equals "name2")
func (s *SelectSuite) Test_SelectString_WhereMulti() {
	sb := fsb.Select().
		From(fsb.Table("users")).
		Where(fsb.Eq("name", "test").
			OR(
				fsb.Eq("id", 1).AND(fsb.Eq("id", 2)),
			).
			AND(
				fsb.Eq("login_id", "test2").AND(fsb.Eq("name", "name2")),
			),
		)

	sql, err := sb.ToSQL()

	assert.Equal(
		s.T(),
		"SELECT * FROM users WHERE name = 'test' OR (id = 1 AND id = 2) AND login_id = 'test2' AND name = 'name2';",
		sql,
	)
	assert.Nil(s.T(), err)
}

// Test_SelectString_InnerJoin tests the SelectString method in the SelectSuite struct.
func (s *SelectSuite) Test_SelectString_InnerJoin() {
	sb := fsb.Select().
		From(fsb.Table("users")).
		InnerJoin(fsb.Table("tokens"), fsb.Eq("users.id", "1"))

	sql, err := sb.ToSQL()

	assert.Equal(
		s.T(),
		"SELECT * FROM users INNER JOIN tokens ON users.id = '1';",
		sql,
	)
	assert.Nil(s.T(), err)
}

// Test_SelectString_InnerJoinCol tests the SelectString method with InnerJoinCol in the SelectSuite struct.
func (s *SelectSuite) Test_SelectString_InnerJoinCol() {
	token := fsb.Table("tokens")

	sb := fsb.Select().
		From(fsb.Table("users")).
		InnerJoin(token, fsb.Eq("users.id", token.Col("user_id")))

	sql, err := sb.ToSQL()

	assert.Equal(
		s.T(),
		"SELECT * FROM users INNER JOIN tokens ON users.id = tokens.user_id;",
		sql,
	)
	assert.Nil(s.T(), err)
}

// Test_SelectString_InnerJoinTable tests the SelectString method in the SelectSuite struct
func (s *SelectSuite) Test_SelectString_InnerJoinTable() {
	user := fsb.Table("users").As("u")
	token := fsb.Table("tokens").As("t")

	sb := fsb.Select(user.Col("id")).
		From(user).
		InnerJoin(token, fsb.Eq(user.Col("id"), token.Col("user_id")))

	sql, err := sb.ToSQL()

	assert.Equal(
		s.T(),
		"SELECT u.id FROM users AS u INNER JOIN tokens AS t ON u.id = t.user_id;",
		sql,
	)
	assert.Nil(s.T(), err)
}

// Test_SelectString_LeftJoin tests the LeftJoin method in the SelectContainer struct.
// It joins the "users" table with the "tokens" table using the condition "users.id = '1'".
// It then calls the ToSQL method to generate the SQL query string and checks if it matches the expected query.
func (s *SelectSuite) Test_SelectString_LeftJoin() {
	sb := fsb.Select().
		From(fsb.Table("users")).
		LeftJoin(fsb.Table("tokens"), fsb.Eq("users.id", "1"))

	sql, err := sb.ToSQL()

	assert.Equal(
		s.T(),
		"SELECT * FROM users LEFT JOIN tokens ON users.id = '1';",
		sql,
	)
	assert.Nil(s.T(), err)
}

// Test_SelectString_LeftJoinCol tests the SelectString method in the SelectSuite struct
// with a LeftJoin operation on two tables: "users" and "tokens" based on the condition "users.id
func (s *SelectSuite) Test_SelectString_LeftJoinCol() {
	token := fsb.Table("tokens")

	sb := fsb.Select().
		From(fsb.Table("users")).
		LeftJoin(token, fsb.Eq("users.id", token.Col("user_id")))

	sql, err := sb.ToSQL()

	assert.Equal(
		s.T(),
		"SELECT * FROM users LEFT JOIN tokens ON users.id = tokens.user_id;",
		sql,
	)
	assert.Nil(s.T(), err)
}

// Test_SelectString_LeftJoinTable tests the SelectString method in the SelectSuite struct.
func (s *SelectSuite) Test_SelectString_LeftJoinTable() {
	user := fsb.Table("users").As("u")
	token := fsb.Table("tokens").As("t")

	sb := fsb.Select(user.Col("id")).
		From(user).
		LeftJoin(token, fsb.Eq(user.Col("id"), token.Col("user_id")))

	sql, err := sb.ToSQL()

	assert.Equal(
		s.T(),
		"SELECT u.id FROM users AS u LEFT JOIN tokens AS t ON u.id = t.user_id;",
		sql,
	)
	assert.Nil(s.T(), err)
}

// Test_SelectString_RightJoin tests the SelectString method in the SelectSuite struct
// with a right join between the "users" and "tokens" tables on the condition "users.id = '1'".
// It verifies that the generated SQL query is correct and that no errors occur.
// Example usage:
//
//	testSuite := &SelectSuite{}
//	testSuite.Test_SelectString_RightJoin()
//
// Expected Output:
//
//	SELECT * FROM users RIGHT JOIN tokens ON users.id = '1';
func (s *SelectSuite) Test_SelectString_RightJoin() {
	sb := fsb.Select().
		From(fsb.Table("users")).
		RightJoin(fsb.Table("tokens"), fsb.Eq("users.id", "1"))

	sql, err := sb.ToSQL()

	assert.Equal(
		s.T(),
		"SELECT * FROM users RIGHT JOIN tokens ON users.id = '1';",
		sql,
	)
	assert.Nil(s.T(), err)
}

// Test_SelectString_RightJoinCol tests the SelectString method in the SelectSuite struct
// by constructing a SELECT query with a Right Join and asserting that the generated SQL
// matches the expected SQL statement.
// The query joins the "users" table with the "tokens" table on the condition that "users.id"
// is equal to "tokens.user_id".
// The expected SQL statement is "SELECT * FROM users RIGHT JOIN tokens ON users.id = tokens.user_id;".
// The generated SQL statement is compared with the expected SQL statement and any error in the process
// is checked to be nil using the assert package's Equal and Nil functions, respectively.
func (s *SelectSuite) Test_SelectString_RightJoinCol() {
	token := fsb.Table("tokens")

	sb := fsb.Select().
		From(fsb.Table("users")).
		RightJoin(token, fsb.Eq("users.id", token.Col("user_id")))

	sql, err := sb.ToSQL()

	assert.Equal(
		s.T(),
		"SELECT * FROM users RIGHT JOIN tokens ON users.id = tokens.user_id;",
		sql,
	)
	assert.Nil(s.T(), err)
}

// Test_SelectString_RightJoinTable tests the SelectString method in the SelectSuite struct.
// It tests the creation of a SELECT statement with a right join on two tables: users and tokens.
// The SELECT statement selects the id column from the users table.
// The right join is performed using the user_id column of the tokens table and the id column of the users table.
// The resulting SQL statement is compared to the expected SQL statement:
// "SELECT u.id FROM users AS u RIGHT JOIN tokens AS t ON u.id = t.user_id;".
// The method also checks for any error that may occur during the creation of the SQL statement.
func (s *SelectSuite) Test_SelectString_RightJoinTable() {
	user := fsb.Table("users").As("u")
	token := fsb.Table("tokens").As("t")

	sb := fsb.Select(user.Col("id")).
		From(user).
		RightJoin(token, fsb.Eq(user.Col("id"), token.Col("user_id")))

	sql, err := sb.ToSQL()

	assert.Equal(
		s.T(),
		"SELECT u.id FROM users AS u RIGHT JOIN tokens AS t ON u.id = t.user_id;",
		sql,
	)
	assert.Nil(s.T(), err)
}

// Test_SelectString_FullJoin tests the SelectString method in the SelectSuite struct
// by creating a SelectBuilder that performs a full join between the "users" table and
// the "tokens" table using the condition "users.id = '1'". It then calls the ToSQL
// method to generate the SQL query.
//
// It asserts that the generated SQL query is equal to the expected query
// "SELECT * FROM users FULL JOIN tokens ON users.id = '1';" and that the error is nil.
func (s *SelectSuite) Test_SelectString_FullJoin() {
	sb := fsb.Select().
		From(fsb.Table("users")).
		FullJoin(fsb.Table("tokens"), fsb.Eq("users.id", "1"))

	sql, err := sb.ToSQL()

	assert.Equal(
		s.T(),
		"SELECT * FROM users FULL JOIN tokens ON users.id = '1';",
		sql,
	)
	assert.Nil(s.T(), err)
}

// Test_SelectString_FullJoinCol tests the SelectString method in the SelectSuite struct
// It creates a SQL SELECT statement with a FULL JOIN using the fsb.Table function,
// and asserts that the generated SQL matches the expected SQL string.
// If the SQL generation returns an error, it asserts that the error is nil.
func (s *SelectSuite) Test_SelectString_FullJoinCol() {
	token := fsb.Table("tokens")

	sb := fsb.Select().
		From(fsb.Table("users")).
		FullJoin(token, fsb.Eq("users.id", token.Col("user_id")))

	sql, err := sb.ToSQL()

	assert.Equal(
		s.T(),
		"SELECT * FROM users FULL JOIN tokens ON users.id = tokens.user_id;",
		sql,
	)
	assert.Nil(s.T(), err)
}

// Test_SelectString_FullJoinTable tests the SelectString method in the SelectSuite struct
// It tests the scenario
// where we have a full join between two tables and select only the "id" column from the "users" table
// The generated SQL statement should be "SELECT u.id FROM users AS u FULL JOIN tokens AS t ON u.id = t.user_id;"
// The test verifies that the expected SQL statement is generated and no error occurs during the process.
func (s *SelectSuite) Test_SelectString_FullJoinTable() {
	user := fsb.Table("users").As("u")
	token := fsb.Table("tokens").As("t")

	sb := fsb.Select(user.Col("id")).
		From(user).
		FullJoin(token, fsb.Eq(user.Col("id"), token.Col("user_id")))

	sql, err := sb.ToSQL()

	assert.Equal(
		s.T(),
		"SELECT u.id FROM users AS u FULL JOIN tokens AS t ON u.id = t.user_id;",
		sql,
	)
	assert.Nil(s.T(), err)
}

// Test_SelectString_CrossJoin tests the SelectString method in the SelectSuite struct
// It creates a select query, performs a cross join of two tables (users and tokens),
// and checks if the generated SQL statement is correct.
func (s *SelectSuite) Test_SelectString_CrossJoin() {
	sb := fsb.Select().
		From(fsb.Table("users")).
		CrossJoin(fsb.Table("tokens"))

	sql, err := sb.ToSQL()

	assert.Equal(
		s.T(),
		"SELECT * FROM users CROSS JOIN tokens;",
		sql,
	)
	assert.Nil(s.T(), err)
}

// Test_SelectString_CrossJoinTable tests the SelectString method in the SelectSuite struct
// by performing a cross join between the "users" and "tokens" tables.
func (s *SelectSuite) Test_SelectString_CrossJoinTable() {
	user := fsb.Table("users").As("u")
	token := fsb.Table("tokens").As("t")

	sb := fsb.Select(user.Col("id")).
		From(user).
		CrossJoin(token)

	sql, err := sb.ToSQL()

	assert.Equal(
		s.T(),
		"SELECT u.id FROM users AS u CROSS JOIN tokens AS t;",
		sql,
	)
	assert.Nil(s.T(), err)
}

// Test_SelectString_OrderString tests the SelectString method in the SelectSuite struct
// with an explicit OrderString.
func (s *SelectSuite) Test_SelectString_OrderString() {
	user := fsb.Table("users").As("u")

	sb := fsb.Select(user.Col("id")).
		From(user).Order("id")

	sql, err := sb.ToSQL()

	assert.Equal(
		s.T(),
		"SELECT u.id FROM users AS u ORDER BY id ASC;",
		sql,
	)
	assert.Nil(s.T(), err)
}

// Test_SelectString_OrderColumn tests the SelectString method in the SelectSuite struct
// by selecting the "id" column from the "users" table, ordering the results by the "id" column in ascending order.
// It then asserts that the generated SQL statement matches the expected SQL statement.
// It also checks that there are no errors returned from the ToSQL method.
func (s *SelectSuite) Test_SelectString_OrderColumn() {
	user := fsb.Table("users").As("u")

	sb := fsb.Select(user.Col("id")).
		From(user).Order(user.Col("id"))

	sql, err := sb.ToSQL()

	assert.Equal(
		s.T(),
		"SELECT u.id FROM users AS u ORDER BY u.id ASC;",
		sql,
	)
	assert.Nil(s.T(), err)
}

// Test_SelectString_OrderASC tests the SelectString method in the SelectSuite struct.
// It checks if the SELECT statement generated by the SelectString method is correct
// when ordering the results in ascending order.
func (s *SelectSuite) Test_SelectString_OrderASC() {
	user := fsb.Table("users").As("u")

	sb := fsb.Select(user.Col("id")).
		From(user).Order("id").DESC().ASC()

	sql, err := sb.ToSQL()

	assert.Equal(
		s.T(),
		"SELECT u.id FROM users AS u ORDER BY id ASC;",
		sql,
	)
	assert.Nil(s.T(), err)
}

// Test_SelectString_OrderDESC tests the SelectString method in the SelectSuite struct
// with the order by clause in descending order.
// It selects the 'id' column from the 'users' table and orders the result by 'id' in descending order.
// The expected SQL query is "SELECT u.id FROM users AS u ORDER BY id DESC;"
// It asserts that the generated SQL query matches the expected query and that no errors occur.
func (s *SelectSuite) Test_SelectString_OrderDESC() {
	user := fsb.Table("users").As("u")

	sb := fsb.Select(user.Col("id")).
		From(user).Order("id").DESC()

	sql, err := sb.ToSQL()

	assert.Equal(
		s.T(),
		"SELECT u.id FROM users AS u ORDER BY id DESC;",
		sql,
	)
	assert.Nil(s.T(), err)
}

// Test_SelectString_OrderASC tests the SelectString method in the SelectSuite struct.
// It checks if the SELECT statement generated by the SelectString method is correct
// when ordering the results in ascending order.
func (s *SelectSuite) Test_SelectString_OrderMulti() {
	user := fsb.Table("users").As("u")

	sb := fsb.Select(user.Col("id")).
		From(user).Order("id").DESC().Order("user_id").ASC()

	sql, err := sb.ToSQL()

	assert.Equal(
		s.T(),
		"SELECT u.id FROM users AS u ORDER BY id DESC, user_id ASC;",
		sql,
	)
	assert.Nil(s.T(), err)
}

// Test_SelectString_OrderAString tests the SelectString method in the SelectSuite struct.
func (s *SelectSuite) Test_SelectString_OrderAString() {
	user := fsb.Table("users").As("u")

	sb := fsb.Select(user.Col("id")).
		From(user).OrderA("id")

	sql, err := sb.ToSQL()

	assert.Equal(
		s.T(),
		"SELECT u.id FROM users AS u ORDER BY id ASC;",
		sql,
	)
	assert.Nil(s.T(), err)
}

// Test_SelectString_OrderAColumn tests the SelectString_OrderAColumn method in the SelectSuite struct.
func (s *SelectSuite) Test_SelectString_OrderAColumn() {
	user := fsb.Table("users").As("u")

	sb := fsb.Select(user.Col("id")).
		From(user).OrderA(user.Col("id"))

	sql, err := sb.ToSQL()

	assert.Equal(
		s.T(),
		"SELECT u.id FROM users AS u ORDER BY u.id ASC;",
		sql,
	)
	assert.Nil(s.T(), err)
}

// Test_SelectString_OrderDeString tests the SelectString_OrderDeString method in the SelectSuite struct.
func (s *SelectSuite) Test_SelectString_OrderDeString() {
	user := fsb.Table("users").As("u")

	sb := fsb.Select(user.Col("id")).
		From(user).OrderDe("id")

	sql, err := sb.ToSQL()

	assert.Equal(
		s.T(),
		"SELECT u.id FROM users AS u ORDER BY id DESC;",
		sql,
	)
	assert.Nil(s.T(), err)
}

// Test_SelectString_OrderDeColumn tests the SelectString method in the SelectSuite struct.
func (s *SelectSuite) Test_SelectString_OrderDeColumn() {
	user := fsb.Table("users").As("u")

	sb := fsb.Select(user.Col("id")).
		From(user).OrderDe(user.Col("id"))

	sql, err := sb.ToSQL()

	assert.Equal(
		s.T(),
		"SELECT u.id FROM users AS u ORDER BY u.id DESC;",
		sql,
	)
	assert.Nil(s.T(), err)
}

// Test_SelectString_Limit tests the SelectString method in the SelectSuite struct.
func (s *SelectSuite) Test_SelectString_Limit() {
	user := fsb.Table("users").As("u")

	sb := fsb.Select(user.Col("id")).
		From(user).Limit(5)

	sql, err := sb.ToSQL()

	assert.Equal(
		s.T(),
		"SELECT u.id FROM users AS u LIMIT 5;",
		sql,
	)
	assert.Nil(s.T(), err)
}

// Test_SelectString_Offset tests the SelectString method in the SelectSuite struct.
func (s *SelectSuite) Test_SelectString_Offset() {
	user := fsb.Table("users").As("u")

	sb := fsb.Select(user.Col("id")).
		From(user).Offset(5)

	sql, err := sb.ToSQL()

	assert.Equal(
		s.T(),
		"SELECT u.id FROM users AS u OFFSET 5;",
		sql,
	)
	assert.Nil(s.T(), err)
}

// Test_SelectString_GroupByString tests the SelectString method in the SelectSuite struct.
func (s *SelectSuite) Test_SelectString_GroupByString() {
	user := fsb.Table("users").As("u")

	sb := fsb.Select(user.Col("id")).
		From(user).GroupBy("id")

	sql, err := sb.ToSQL()

	assert.Equal(
		s.T(),
		"SELECT u.id FROM users AS u GROUP BY id;",
		sql,
	)
	assert.Nil(s.T(), err)
}

// Test_SelectString_GroupByMultiString tests the SelectString_GroupByMultiString method in the SelectSuite struct.
func (s *SelectSuite) Test_SelectString_GroupByMultiString() {
	user := fsb.Table("users").As("u")

	sb := fsb.Select(user.Col("id"), user.Col("name")).
		From(user).GroupBy("id", "name")

	sql, err := sb.ToSQL()

	assert.Equal(
		s.T(),
		"SELECT u.id, u.name FROM users AS u GROUP BY id, name;",
		sql,
	)
	assert.Nil(s.T(), err)
}

// Test_SelectString_GroupByMultiColumn tests the SelectString method in the SelectSuite struct.
func (s *SelectSuite) Test_SelectString_GroupByMultiColumn() {
	user := fsb.Table("users").As("u")

	sb := fsb.Select(user.Col("id"), user.Col("name")).
		From(user).GroupBy(user.Col("id"), user.Col("name"))

	sql, err := sb.ToSQL()

	assert.Equal(
		s.T(),
		"SELECT u.id, u.name FROM users AS u GROUP BY u.id, u.name;",
		sql,
	)
	assert.Nil(s.T(), err)
}

// Test_SelectString_GroupByMulti tests the SelectString_GroupByMulti method in the SelectSuite struct.
func (s *SelectSuite) Test_SelectString_GroupByMulti() {
	user := fsb.Table("users").As("u")

	sb := fsb.Select(user.Col("id"), user.Col("name")).
		From(user).GroupBy(user.Col("id"), "name")

	sql, err := sb.ToSQL()

	assert.Equal(
		s.T(),
		"SELECT u.id, u.name FROM users AS u GROUP BY u.id, name;",
		sql,
	)
	assert.Nil(s.T(), err)
}

// Test_SelectString_GroupByColumn tests the SelectString method in the SelectSuite struct.
func (s *SelectSuite) Test_SelectString_GroupByColumn() {
	user := fsb.Table("users").As("u")

	sb := fsb.Select(user.Col("id")).
		From(user).GroupBy(user.Col("id"))

	sql, err := sb.ToSQL()

	assert.Equal(
		s.T(),
		"SELECT u.id FROM users AS u GROUP BY u.id;",
		sql,
	)
	assert.Nil(s.T(), err)
}

// Test_SelectString_HavingInt tests the SelectString method in the SelectSuite struct
// with a HAVING clause that compares the "id" column to the integer 1.
func (s *SelectSuite) Test_SelectString_HavingInt() {
	sb := fsb.Select().
		From(fsb.Table("users")).
		Having(fsb.Eq("id", 1))

	sql, err := sb.ToSQL()

	assert.Equal(s.T(), "SELECT * FROM users HAVING id = 1;", sql)
	assert.Nil(s.T(), err)
}

// Test_SelectString_HavingString tests the SelectString method in the SelectSuite struct
func (s *SelectSuite) Test_SelectString_HavingString() {
	sb := fsb.Select().
		From(fsb.Table("users")).
		Having(fsb.Eq("name", "test"))

	sql, err := sb.ToSQL()

	assert.Equal(s.T(), "SELECT * FROM users HAVING name = 'test';", sql)
	assert.Nil(s.T(), err)
}

func (s *SelectSuite) Test_SelectString_HavingAND() {
	sb := fsb.Select().
		From(fsb.Table("users")).
		Having(fsb.Eq("name", "test").AND(fsb.Eq("id", 1)))

	sql, err := sb.ToSQL()

	assert.Equal(s.T(), "SELECT * FROM users HAVING name = 'test' AND id = 1;", sql)
	assert.Nil(s.T(), err)
}

// Test_SelectString_HavingOR tests the SelectString method in the SelectSuite struct.
func (s *SelectSuite) Test_SelectString_HavingOR() {
	sb := fsb.Select().
		From(fsb.Table("users")).
		Having(fsb.Eq("name", "test").OR(fsb.Eq("id", 1)))

	sql, err := sb.ToSQL()

	assert.Equal(s.T(), "SELECT * FROM users HAVING name = 'test' OR id = 1;", sql)
	assert.Nil(s.T(), err)
}

// Test_SelectString_HavingAndOr tests the SelectString method in the SelectSuite struct
// with the Having and Or conditions.
func (s *SelectSuite) Test_SelectString_HavingAndOr() {
	sb := fsb.Select().
		From(fsb.Table("users")).
		Having(fsb.Eq("name", "test").
			AND(
				fsb.Eq("id", 1).OR(fsb.Eq("id", 2)),
			),
		)

	sql, err := sb.ToSQL()

	assert.Equal(s.T(), "SELECT * FROM users HAVING name = 'test' AND (id = 1 OR id = 2);", sql)
	assert.Nil(s.T(), err)
}

// Test_SelectString_HavingOrAnd tests the SelectString method in the SelectSuite struct,
// specifically testing the functionality of adding having conditions with OR and AND clauses.
func (s *SelectSuite) Test_SelectString_HavingOrAnd() {
	sb := fsb.Select().
		From(fsb.Table("users")).
		Having(fsb.Eq("name", "test").
			OR(
				fsb.Eq("id", 1).AND(fsb.Eq("id", 2)),
			),
		)

	sql, err := sb.ToSQL()

	assert.Equal(s.T(), "SELECT * FROM users HAVING name = 'test' OR (id = 1 AND id = 2);", sql)
	assert.Nil(s.T(), err)
}

// Test_SelectString_HavingMulti tests the SelectString method in the SelectSuite struct with multiple conditions in the HAVING clause.
func (s *SelectSuite) Test_SelectString_HavingMulti() {
	sb := fsb.Select().
		From(fsb.Table("users")).
		Having(fsb.Eq("name", "test").
			OR(
				fsb.Eq("id", 1).AND(fsb.Eq("id", 2)),
			).
			AND(
				fsb.Eq("login_id", "test2").AND(fsb.Eq("name", "name2")),
			),
		)

	sql, err := sb.ToSQL()

	assert.Equal(
		s.T(),
		"SELECT * FROM users HAVING name = 'test' OR (id = 1 AND id = 2) AND login_id = 'test2' AND name = 'name2';",
		sql,
	)
	assert.Nil(s.T(), err)
}

func TestSelectSuite(t *testing.T) {
	suite.Run(t, new(SelectSuite))
}

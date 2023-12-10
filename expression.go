package fsb

import (
	"fmt"
	"strconv"
	"strings"
)

// Expression
// AND combines the current expression with another expression using the AND operator.
// If the other expression contains the OR operator, the current expression will be enclosed in brackets.
type Expression struct {
	condition  string
	bracketFlg bool
}

// Eq is a function that creates an Expression with a specific condition based on the target and comparison value.
// It can handle string and int comparison values.
// The condition is built using the fmt.Sprintf function to format the target and comparison value appropriately.
// The function returns a pointer to an Expression struct initialized with the condition.
func Eq(target string, comp interface{}) *Expression {
	cond := createCondition(target, comp, "=")

	return &Expression{
		condition: cond,
	}
}

// Neq is a function that creates an Expression with a specific condition based on the target and comparison value.
// It can handle string and int comparison values.
// The condition is built using the fmt.Sprintf function to format the target and comparison value appropriately with the "!=" operator.
// The function returns a pointer to an Expression struct initialized with the condition.
func Neq(target string, comp interface{}) *Expression {
	cond := createCondition(target, comp, "!=")

	return &Expression{
		condition: cond,
	}
}

// Gt is a function that creates an Expression with a specific condition based on the target and comparison value,
// where the target is greater than the comparison value.
// It can handle string and int comparison values.
// The condition is built using the fmt.Sprintf function to format the target and comparison value appropriately.
// The function returns a pointer to an Expression struct initialized with the condition.
func Gt(target string, comp interface{}) *Expression {
	cond := createCondition(target, comp, ">")

	return &Expression{
		condition: cond,
	}
}

// Gte is a function that creates an Expression with a specific condition based on the target and comparison value.
// It can handle string and int comparison values.
// The condition is built using the fmt.Sprintf function to format the target and comparison value appropriately,
// with ">=" as the comparison operator.
// The function returns a pointer to an Expression struct initialized with the condition.
func Gte(target string, comp interface{}) *Expression {
	cond := createCondition(target, comp, ">=")

	return &Expression{
		condition: cond,
	}
}

// Lt is a function that creates an Expression with a specific condition based on the target and comparison value.
// It can handle string and int comparison values.
// The condition is built using the fmt.Sprintf function to format the target and comparison value appropriately,
// with a less than (<) symbol.
// The function returns a pointer to an Expression struct initialized with the condition.
func Lt(target string, comp interface{}) *Expression {
	cond := createCondition(target, comp, "<")

	return &Expression{
		condition: cond,
	}
}

// Lte is a function that creates an Expression with a specific condition based on the target and comparison value.
// It can handle string and int comparison values.
// The condition is built using the fmt.Sprintf function to format the target and comparison value appropriately.
// The function returns a pointer to an Expression struct initialized with the condition.
func Lte(target string, comp interface{}) *Expression {
	cond := createCondition(target, comp, "<=")

	return &Expression{
		condition: cond,
	}
}

// Like is a function that creates an Expression with a specific condition based
// on the target and comparison value using the "LIKE" operator.
// It handles only string comparison values.
// The condition is built using the fmt.Sprintf function to format the target and comparison value appropriately.
// The function returns a pointer to an Expression struct initialized with the condition.
func Like(target, comp string) *Expression {
	cond := createCondition(target, comp, "LIKE")

	return &Expression{
		condition: cond,
	}
}

// Nlike is a function that creates an Expression struct with a specific condition based
// on the target and comparison value.
// It uses the "NOT LIKE" sign to build the condition.
// The function returns a pointer to the Expression struct initialized with the condition.
func Nlike(target, comp string) *Expression {
	cond := createCondition(target, comp, "NOT LIKE")

	return &Expression{
		condition: cond,
	}
}

// Pm is a function that creates an Expression with a condition using the "LIKE" operator.
// It takes a target string and a comparison value as arguments.
// The comparison value is converted to a SQL like prefix pattern using the sqlLikePrefixPattern function.
// The createCondition function is used to create the condition string with the target,
// converted comparison value, and "LIKE" operator.
// The function returns a pointer to an Expression struct initialized with the condition.
func Pm(target string, comp interface{}) *Expression {
	cond := createCondition(target, sqlLikePrefixPattern(comp), "LIKE")

	return &Expression{
		condition: cond,
	}
}

// Npm is a function that creates an Expression with a specific condition based on the target and comparison value.
// It uses the createCondition function to build the condition using the target,
// the comparison value converted into a SQL like prefix pattern, and the "NOT LIKE" sign.
// The function returns a pointer to an Expression struct initialized with the condition.
func Npm(target string, comp interface{}) *Expression {
	cond := createCondition(target, sqlLikePrefixPattern(comp), "NOT LIKE")

	return &Expression{
		condition: cond,
	}
}

// Sm is a function that creates an Expression with a specific condition based on the target and comparison value.
// It uses the createCondition function to build the condition string using the target,
// a modified comparison value obtained from the sqlLikeSuffixPattern function, and the "LIKE"
func Sm(target string, comp interface{}) *Expression {
	cond := createCondition(target, sqlLikeSuffixPattern(comp), "LIKE")

	return &Expression{
		condition: cond,
	}
}

// Nsm is a function that creates an Expression with a specific "NOT LIKE" condition
// based on the target and comparison value.
// It can handle string and int comparison values.
// The comparison value is converted to a SQL LIKE suffix pattern using the sqlLikeSuffixPattern function.
// The condition is built using the createCondition function with the target,
// SQL LIKE suffix pattern, and "NOT LIKE" sign.
// The function returns a pointer to an Expression struct initialized with the condition.
func Nsm(target string, comp interface{}) *Expression {
	cond := createCondition(target, sqlLikeSuffixPattern(comp), "NOT LIKE")

	return &Expression{
		condition: cond,
	}
}

// Psm is a function that creates an Expression with a condition using the LIKE operator.
// It takes a target string and a comp interface{} as arguments.
// The function first converts the comp value into a SQL pattern with a prefix
// and suffix using the sqlLikePrefixPattern and sqlLikeSuffixPattern functions.
// Then, it calls the createCondition function to create the condition string using the target,
// the SQL pattern, and the "LIKE" sign.
// The function returns a pointer to an Expression struct initialized with the condition.
func Psm(target string, comp interface{}) *Expression {
	cond := createCondition(target, sqlLikePrefixPattern(sqlLikeSuffixPattern(comp)), "LIKE")

	return &Expression{
		condition: cond,
	}
}

// Npsm is a function that creates an Expression with a specific condition based on the target and comparison value.
// It can handle string and int comparison values.
// The comp value is converted to a SQL LIKE pattern with a prefix
// and suffix using the sqlLikePrefixPattern and sqlLikeSuffixPattern functions.
// The condition is built using the createCondition function with the target, converted comp value, and "NOT LIKE" sign.
// The function returns a pointer to an Expression struct initialized with the condition.
func Npsm(target string, comp interface{}) *Expression {
	cond := createCondition(target, sqlLikePrefixPattern(sqlLikeSuffixPattern(comp)), "NOT LIKE")

	return &Expression{
		condition: cond,
	}
}

// Between is a function that creates an Expression
// with a specific condition based on the target, start, and end values.
// It can handle string and int start values and expects the end value to have the same type as the start value.
// The condition is built using the fmt.Sprintf function to format the target and values appropriately.
// The function returns a pointer to an Expression struct initialized with the condition.
func Between(target string, start, end interface{}) *Expression {
	var cond string
	switch v := start.(type) {
	case string:
		cond = fmt.Sprintf("%s BETWEEN '%s' TO '%s'", target, v, end.(string))
	case int:
		cond = fmt.Sprintf("%s BETWEEN %d TO %d", target, v, end.(int))
	}

	return &Expression{
		condition: cond,
	}
}

// createCondition is a function that takes a target string, comparison value, and sign string
// and returns a string representing the condition for the expression.
// It supports string and int comparison values and formats the condition using fmt.Sprintf.
// The returned condition string includes the target, sign, and comparison value appropriately formatted.
func createCondition(target string, comp interface{}, sign string) string {
	switch v := comp.(type) {
	case string:
		return fmt.Sprintf("%s %s '%s'", target, sign, v)
	case int:
		return fmt.Sprintf("%s %s %d", target, sign, v)
	default:
		return ""
	}
}

// toSqlLikePattern is a function that converts a comparison value to a SQL LIKE pattern.
// It takes an interface{} as an argument and returns a string.
// If the comparison value is a string, it appends '%' at the end.
// If the comparison value is an int, it converts it to a string and appends '%' at the end.
// For any other type, it returns an empty string.
func sqlLikePrefixPattern(comp interface{}) string {
	switch v := comp.(type) {
	case string:
		return v + "%"
	case int:
		return strconv.Itoa(v) + "%"
	default:
		return ""
	}
}

// sqlLikeSuffixPattern is a function that takes a comparison value as input
// and returns a string representation of a SQL "LIKE" pattern with a suffix wildcard.
// It can handle comparison values of type string and int.
// If the comparison value is a string, the function prepends a "%" character to it.
// If the comparison value is an int,
// the function converts it to a string using strconv.Itoa and prepends a "%" character to it.
// If the comparison value is of any other type, the function returns an empty string.
// The function is primarily used in conjunction with the createCondition function to generate a SQL
// WHERE clause condition for a "LIKE" comparison.
// Example usages of sqlLikeSuffixPattern can be seen in the Sm and Nsm functions.
func sqlLikeSuffixPattern(comp interface{}) string {
	switch v := comp.(type) {
	case string:
		return "%" + v
	case int:
		return "%" + strconv.Itoa(v)
	default:
		return ""
	}
}

// AND sets the condition of the Expression object with the logical AND operator.
// If the exp.condition contains the string "OR", it appends the condition in brackets with the AND operator.
// Otherwise, it appends the condition with the AND operator without brackets.
//
// Example usage:
//
//	exp := &Expression{condition: "name = 'test'"}
//	newExp := exp.AND(&Expression{condition: "id = 1"})
//	// Output: newExp.condition = "name = 'test' AND id = 1"
//
//	exp := &Expression{condition: "name = 'test'"}
//	newExp := exp.AND(&Expression{condition: "id = 1 OR id = 2"})
//	// Output: newExp.condition = "name = 'test' AND (id = 1 OR id = 2)"
//
//	exp := &Expression{condition: "name = 'test'"}
//	newExp := exp.AND(&Expression{condition: "id = 1 AND id = 2"})
//	// Output: newExp.condition = "name = 'test' AND (id = 1 AND id = 2)"
func (e *Expression) AND(exp *Expression) *Expression {
	if strings.Contains(exp.condition, "OR") {
		e.condition = fmt.Sprintf("%s AND (%s)", e.condition, exp.condition)
	} else {
		e.condition = fmt.Sprintf("%s AND %s", e.condition, exp.condition)
	}

	return e
}

// OR sets the condition of the Expression object with the logical OR operator.
// If the exp.condition contains the string "AND", it appends the condition in brackets with the OR operator.
// Otherwise, it appends the condition with the OR operator without brackets.
//
// It also sets the bracketFlg to true, indicating that the condition contains brackets.
//
// Example usage:
//
//	exp := &Expression{condition: "name = 'test'"}
//	newExp := exp.OR(&Expression{condition: "id = 1"})
//	// Output: newExp.condition = "name = 'test' OR id = 1"
//
//	exp := &Expression{condition: "name = 'test'"}
//	newExp := exp.OR(&Expression{condition: "id = 1 AND id = 2"})
//	// Output: newExp.condition = "name = 'test' OR (id = 1 AND id = 2)"
//
//	exp := &Expression{condition: "name = 'test'"}
//	newExp := exp.OR(&Expression{condition: "id = 1 OR id = 2"})
//	// Output: newExp.condition = "name = 'test' OR (id = 1 OR id = 2)"
func (e *Expression) OR(exp *Expression) *Expression {
	if strings.Contains(exp.condition, "AND") {
		e.condition = fmt.Sprintf("%s OR (%s)", e.condition, exp.condition)
	} else {
		e.condition = fmt.Sprintf("%s OR %s", e.condition, exp.condition)
	}

	e.bracketFlg = true

	return e
}

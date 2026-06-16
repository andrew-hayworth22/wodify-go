package query

import (
	"errors"
	"fmt"
	"strings"
)

// ErrReservedChar is returned when a query value contains a reserved character
var ErrReservedChar = errors.New("query value contains a reserved character")

const (
	baseReserved = "'|;" // reserved characters in all clauses
	listReserved = "{}," // reserved characters in list clauses (in/not_in)
)

// clause represents a query clause
type clause struct {
	field    string
	operator operator
	values   []string
}

// newClause creates a new query clause
func newClause(field string, operator operator, values ...any) (clause, error) {
	reservedCharacters := baseReserved
	if operator == in || operator == notIn {
		reservedCharacters += listReserved
	}

	valuesStrs := make([]string, len(values))
	for i, v := range values {
		switch val := v.(type) {
		case string:
			if i := strings.IndexAny(val, reservedCharacters); i >= 0 {
				return clause{}, fmt.Errorf("%w: field %q value %q contains %q", ErrReservedChar, field, val, val[i])
			}
			valuesStrs[i] = fmt.Sprintf("'%s'", val)
		default:
			valuesStrs[i] = fmt.Sprintf("%v", val)
		}
	}

	return clause{
		field:    field,
		operator: operator,
		values:   valuesStrs,
	}, nil
}

// encode converts a query clause into a string
func (c *clause) encode() string {
	switch c.operator {
	case null, notNull:
		return fmt.Sprintf("%s|%s", c.field, c.operator)
	case between:
		return fmt.Sprintf("%s|%s|%s|%s", c.field, c.operator, c.values[0], c.values[1])
	case in, notIn:
		return fmt.Sprintf("%s|%s|{%s}", c.field, c.operator, strings.Join(c.values, ","))
	default:
		return fmt.Sprintf("%s|%s|%s", c.field, c.operator, c.values[0])
	}
}

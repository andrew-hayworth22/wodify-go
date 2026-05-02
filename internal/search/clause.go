package search

import (
	"fmt"
	"strings"
)

// clause represents a search clause
type clause struct {
	field    string
	operator operator
	values   []string
}

// newClause creates a new search clause
func newClause(field string, operator operator, values ...any) clause {
	valuesStrs := make([]string, len(values))
	for i, v := range values {
		switch v.(type) {
		case string:
			valuesStrs[i] = fmt.Sprintf("'%v'", v)
		default:
			valuesStrs[i] = fmt.Sprintf("%v", v)
		}
	}

	return clause{
		field:    field,
		operator: operator,
		values:   valuesStrs,
	}
}

// encode converts a search clause into a string
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

package search

// operator represents a supported search operator
type operator string

const (
	lessThan           operator = "lt"
	lessThanEqualTo    operator = "lte"
	greaterThan        operator = "gt"
	greaterThanEqualTo operator = "gte"
	equal              operator = "eq"
	notEqual           operator = "neq"

	between    operator = "between"
	in         operator = "in"
	notIn      operator = "not_in"
	null       operator = "is_null"
	notNull    operator = "not_null"
	contains   operator = "contains"
	startsWith operator = "starts_with"
	endsWith   operator = "ends_with"
)

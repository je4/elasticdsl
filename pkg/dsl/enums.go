package dsl

type tOperator string

const (
	OperatorOR  tOperator = "OR"
	OperatorAND tOperator = "AND"
)

type tZeroTermsQuery string

const (
	ZeroTermsQueryNONE tZeroTermsQuery = "none"
	ZeroTermsQueryALL  tZeroTermsQuery = "all"
)

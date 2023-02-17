package dsl

type ExistsQuery struct {
	Field string `json:"field"`
}

func (*ExistsQuery) GetQueryName() string { return "exists" }

var (
	_ Query = (*ExistsQuery)(nil)
)

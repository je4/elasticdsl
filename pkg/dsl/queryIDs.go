package dsl

// https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-ids-query.html

type IDsQuery struct {
	IDs []string `json:"ids,omitempty"`
}

func (*IDsQuery) GetQueryName() string { return "ids" }

var (
	_ BaseQuery = (*IDsQuery)(nil)
)

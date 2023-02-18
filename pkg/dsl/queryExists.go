package dsl

import "encoding/json"

// ExistsQuery results a combined_field query struct
//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-exists-query.html
type ExistsQuery func(field string, o ...func(all *tExistsQuery)) BaseQuery

func NewExistsQuery() ExistsQuery {
	return func(field string, o ...func(all *tExistsQuery)) BaseQuery {
		var r = &tExistsQuery{
			Field: field,
		}
		for _, f := range o {
			f(r)
		}
		return r
	}
}

type tExistsQuery struct {
	Field string `json:"field"`
}

func (*tExistsQuery) GetQueryName() string { return "exists" }

func (q *tExistsQuery) MarshalJSON() ([]byte, error) {
	type _tExistsQuery tExistsQuery
	return json.Marshal(map[string]*_tExistsQuery{
		q.GetQueryName(): (*_tExistsQuery)(q),
	})
}

var (
	_ BaseQuery = (*tExistsQuery)(nil)
)

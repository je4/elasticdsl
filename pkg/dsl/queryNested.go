package dsl

import "encoding/json"

// NestedQuery return a nested body
//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-nested-query.html
type NestedQuery func(path string, query *tQuery, o ...func(all *tNestedQuery)) BaseQuery

func (nq *NestedQuery) WithScoreMode(mode ScoreMode) func(all *tNestedQuery) {
	return func(all *tNestedQuery) {
		all.ScoreMode = mode
	}
}

func (nq *NestedQuery) WithIgnoreUnmapped(ignoreUnmapped bool) func(all *tNestedQuery) {
	return func(all *tNestedQuery) {
		all.IgnoreUnmapped = ignoreUnmapped
	}
}

func NewNestedQuery() NestedQuery {
	return func(path string, query *tQuery, o ...func(all *tNestedQuery)) BaseQuery {
		var r = &tNestedQuery{
			Path:  path,
			Query: query,
		}
		for _, f := range o {
			f(r)
		}
		return r
	}
}

type tNestedQuery struct {
	Path           string    `json:"path"`
	Query          *tQuery   `json:"query"`
	ScoreMode      ScoreMode `json:"score_mode,omitempty"`
	IgnoreUnmapped bool      `json:"ignore_unmapped,omitempty"` // default: false
}

func (*tNestedQuery) GetQueryName() string { return "nested" }

func (q *tNestedQuery) MarshalJSON() ([]byte, error) {
	type _tNestedQuery tNestedQuery
	return json.Marshal(map[string]*_tNestedQuery{
		q.GetQueryName(): (*_tNestedQuery)(q),
	})
}

var (
	_ BaseQuery = (*tNestedQuery)(nil)
)

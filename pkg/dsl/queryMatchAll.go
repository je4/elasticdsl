package dsl

import "encoding/json"

// MatchAllQuery results a match_all query struct
//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-match-all-query.html
type MatchAllQuery func(o ...func(all *tMatchAllQuery)) BaseQuery

func (qma *MatchAllQuery) WithBoost(boost float64) func(all *tMatchAllQuery) {
	return func(all *tMatchAllQuery) {
		all.Boost = boost
	}
}

func NewMatchAllQuery() MatchAllQuery {
	return func(o ...func(*tMatchAllQuery)) BaseQuery {
		var r = &tMatchAllQuery{}
		for _, f := range o {
			f(r)
		}
		return r
	}
}

type tMatchAllQuery struct {
	Boost float64 `json:"boost,omitempty"`
}

func (*tMatchAllQuery) GetQueryName() string { return "match_all" }

func (q *tMatchAllQuery) MarshalJSON() ([]byte, error) {
	type _tMatchAllQuery tMatchAllQuery
	return json.Marshal(map[string]*_tMatchAllQuery{
		q.GetQueryName(): (*_tMatchAllQuery)(q),
	})
}

var (
	_ BaseQuery = (*tMatchAllQuery)(nil)
)

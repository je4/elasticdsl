package dsl

import "encoding/json"

// MatchAllQuery results a match_all query struct
//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-match-all-query.html
type MatchNoneQuery func(o ...func(all *tMatchNoneQuery)) BaseQuery

type tMatchNoneQuery struct {
}

func NewMatchNoneQuery() MatchNoneQuery {
	return func(o ...func(query *tMatchNoneQuery)) BaseQuery {
		var r = &tMatchNoneQuery{}
		for _, f := range o {
			f(r)
		}
		return r
	}
}

func (*tMatchNoneQuery) GetQueryName() string { return "match_none" }

func (q *tMatchNoneQuery) MarshalJSON() ([]byte, error) {
	type _tMatchNoneQuery tMatchNoneQuery
	return json.Marshal(map[string]*_tMatchNoneQuery{
		q.GetQueryName(): (*_tMatchNoneQuery)(q),
	})
}

var (
	_ BaseQuery = (*tMatchNoneQuery)(nil)
)

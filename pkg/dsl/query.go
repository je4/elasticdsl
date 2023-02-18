package dsl

import "encoding/json"

// Query returns results a query body.
//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl.html
type Query func(query BaseQuery, o ...func(*tQuery)) *tQuery

func NewQuery() Query {
	return func(query BaseQuery, o ...func(*tQuery)) *tQuery {
		var r = &tQuery{}
		r.Query = query
		for _, f := range o {
			f(r)
		}
		return r
	}
}

type tQuery struct {
	Query BaseQuery `json:"-"`
}

func (q *tQuery) MarshalJSON() ([]byte, error) {
	return json.Marshal(q.Query)
}

func (*tQuery) GetQueryName() string { return "query" }

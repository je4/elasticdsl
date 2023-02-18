package dsl

import "encoding/json"

// BoolQuery results a bool query struct
//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-bool-query.html
type BoolQuery func(o ...func(all *tBoolQuery)) BaseQuery

func (q *BoolQuery) WithMust(bq ...BaseQuery) func(all *tBoolQuery) {
	return func(all *tBoolQuery) {
		all.Must = bq
	}
}

func (q *BoolQuery) WithFilter(bq ...BaseQuery) func(all *tBoolQuery) {
	return func(all *tBoolQuery) {
		all.Filter = bq
	}
}

func (q *BoolQuery) WithShould(bq ...BaseQuery) func(all *tBoolQuery) {
	return func(all *tBoolQuery) {
		all.Should = bq
	}
}

func (q *BoolQuery) WithMustNot(bq ...BaseQuery) func(all *tBoolQuery) {
	return func(all *tBoolQuery) {
		all.MustNot = bq
	}
}

func (q *BoolQuery) WithMinimumShouldMatch(msm *MinimumShouldMatch) func(all *tBoolQuery) {
	return func(all *tBoolQuery) {
		all.MinimumShouldMatch = msm
	}
}

func NewBoolQuery() BoolQuery {
	return func(o ...func(query *tBoolQuery)) BaseQuery {
		var r = &tBoolQuery{}
		for _, f := range o {
			f(r)
		}
		return r
	}
}

type tBoolQuery struct {
	Must               []BaseQuery         `json:"must,omitempty"`
	Filter             []BaseQuery         `json:"filter,omitempty"`
	Should             []BaseQuery         `json:"should,omitempty"`
	MustNot            []BaseQuery         `json:"must_not,omitempty"`
	MinimumShouldMatch *MinimumShouldMatch `json:"minimum_should_match,omitempty"`
}

func (*tBoolQuery) GetQueryName() string { return "bool" }

func (bq *tBoolQuery) MarshalJSON() ([]byte, error) {
	type _tBoolQuery tBoolQuery
	return json.Marshal(map[string]*_tBoolQuery{
		bq.GetQueryName(): (*_tBoolQuery)(bq),
	})
}

var (
	_ BaseQuery = (*tBoolQuery)(nil)
)

package dsl

import "encoding/json"

// Aggs returns an aggs body.
//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/8.6/search-aggregations.html
type Aggs func(o ...func(aggs *tAggs)) *tAggs

func (a *Aggs) WithQueries(qs ...BaseQuery) func(all *tAggs) {
	return func(all *tAggs) {
		all.Queries = qs
	}
}

func (a *Aggs) WithAggs(aggs *tAggs) func(all *tAggs) {
	return func(all *tAggs) {
		all.Aggs = aggs
	}
}

func NewAggs() Aggs {
	return func(o ...func(aggs *tAggs)) *tAggs {
		var r = &tAggs{Queries: []BaseQuery{}}
		for _, f := range o {
			f(r)
		}
		return r
	}
}

type tAggs struct {
	Queries []BaseQuery  `json:"-"`
	Aggs    *tAggs       `json:"-"`
	Nested  *tAggsNested `json:"-"`
}

func (a *tAggs) MarshalJSON() ([]byte, error) {
	var data = map[string]any{}
	if a.Queries != nil {
		for _, q := range a.Queries {
			data[q.GetQueryName()] = q
		}
	}
	if a.Aggs != nil {
		data["aggs"] = a.Aggs
	}
	if a.Nested != nil {
		data["nested"] = a.Nested
	}

	return json.Marshal(data)
}

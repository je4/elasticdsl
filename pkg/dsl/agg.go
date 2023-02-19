package dsl

import "encoding/json"

// Agg returns an agg body.
//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/8.6/search-aggregations.html
type Agg func(name string, query BaseQuery, o ...func(agg *tAgg)) *tAgg

func NewAgg() Agg {
	return func(name string, query BaseQuery, o ...func(agg *tAgg)) *tAgg {
		var r = &tAgg{
			Name:  name,
			Query: query,
		}
		for _, f := range o {
			f(r)
		}
		return r
	}
}

type tAgg struct {
	Name  string    `json:"-"`
	Query BaseQuery `json:"-"`
}

func (agg *tAgg) GetAggName() string { return agg.Name }

func (a *tAgg) MarshalJSON() ([]byte, error) {
	return json.Marshal(a.Query)
}

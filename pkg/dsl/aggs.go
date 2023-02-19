package dsl

import "encoding/json"

// Aggs returns an aggs body.
//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/8.6/search-aggregations.html
type Aggs func(o ...func(aggs *tAggs)) *tAggs

func (a *Aggs) WithAggregations(aggs ...Aggregation) func(all *tAggs) {
	return func(all *tAggs) {
		all.Aggs = aggs
	}
}

func NewAggs() Aggs {
	return func(o ...func(aggs *tAggs)) *tAggs {
		var r = &tAggs{}
		for _, f := range o {
			f(r)
		}
		return r
	}
}

type tAggs struct {
	Aggs []Aggregation `json:"-"`
}

func (*tAggs) GetAggName() string { return "aggs" }

func (a *tAggs) MarshalJSON() ([]byte, error) {
	var data = map[string]any{}
	for _, agg := range a.Aggs {
		data[agg.GetAggName()] = agg
	}

	return json.Marshal(data)
}

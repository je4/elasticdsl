package dsl

import "encoding/json"

// AggFilter returns an AggFilter body.
//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/8.6/search-AggFilterregations.html
type AggFilter func(name string, field string, value string, o ...func(AggFilter *tAggFilter)) *tAggFilter

func (*AggFilter) WithAggs(aggs ...BaseAgg) func(all *tAggFilter) {
	return func(all *tAggFilter) {
		all.Aggs = aggs
	}
}

func NewAggFilter() AggFilter {
	return func(name string, field string, value string, o ...func(AggFilter *tAggFilter)) *tAggFilter {
		var r = &tAggFilter{
			Name:  name,
			Field: field,
			Value: value,
		}
		for _, f := range o {
			f(r)
		}
		return r
	}
}

type tAggFilter struct {
	Name  string    `json:"-"`
	Field string    `json:"-"`
	Value string    `json:"-"`
	Aggs  []BaseAgg `json:"-"`
}

func (a *tAggFilter) GetAggName() string { return a.Name }

func (a *tAggFilter) MarshalJSON() ([]byte, error) {
	var subAggs = map[string]BaseAgg{}
	for _, a := range a.Aggs {
		subAggs[a.GetAggName()] = a
	}
	var data = map[string]any{
		"filter": map[string]any{
			"term": map[string]string{
				a.Field: a.Value,
			},
		},
		"aggs": subAggs,
	}
	return json.Marshal(data)
}

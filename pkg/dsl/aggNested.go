package dsl

import "encoding/json"

// AggNested returns a nested body
//
// See full documentations at https://www.elastic.co/guide/en/elasticsearch/reference/current/search-aggregations-bucket-nested-aggregation.html
type AggNested func(name string, path string, aggs *tAggs, o ...func(aggNestes *tAggNested)) *tAggNested

func NewAggNested() AggNested {
	return func(name string, path string, aggs *tAggs, o ...func(aggNestes *tAggNested)) *tAggNested {
		var r = &tAggNested{
			Name: name,
			Path: path,
			Aggs: aggs,
		}
		for _, f := range o {
			f(r)
		}
		return r
	}
}

type tAggNested struct {
	Name string `json:"-"`
	Path string `json:"-"`
	Aggs *tAggs `json:"-"`
}

func (a *tAggNested) GetAggName() string { return a.Name }

func (an *tAggNested) MarshalJSON() ([]byte, error) {
	data := map[string]any{
		"nested": map[string]any{"path": an.Path},
		"aggs":   an.Aggs,
	}
	return json.Marshal(data)
}

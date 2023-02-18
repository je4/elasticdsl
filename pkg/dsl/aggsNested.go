package dsl

import "encoding/json"

// AggsNested returns a nested body
//
// See full documentations at https://www.elastic.co/guide/en/elasticsearch/reference/current/search-aggregations-bucket-nested-aggregation.html
type AggsNested func(path string, aggs *tAggs, o ...func(aggsNestes *tAggsNested)) *tAggsNested

func NewAggsNested() AggsNested {
	return func(path string, aggs *tAggs, o ...func(aggsNestes *tAggsNested)) *tAggsNested {
		var r = &tAggsNested{
			Path: path,
			Aggs: aggs,
		}
		for _, f := range o {
			f(r)
		}
		return r
	}
}

type tAggsNested struct {
	Path string `json:"path"`
	Aggs *tAggs `json:"aggs"`
}

func (an *tAggsNested) MarshalJSON() ([]byte, error) {
	data := map[string]any{
		"nested": map[string]any{"path": an.Path},
		"aggs":   an.Aggs,
	}
	return json.Marshal(data)
}

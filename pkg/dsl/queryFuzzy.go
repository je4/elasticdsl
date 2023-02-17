package dsl

import "encoding/json"

// https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-fuzzy-query.html

type FuzzyQuery struct {
	Field          string       `json:"-"`
	Value          string       `json:"value"`
	Fuzziness      string       `json:"fuzziness,omitempty"`
	MaxExpansion   int          `json:"max_expansion,omitempty"`  // default: 50
	PrefixLength   int          `json:"prefix_length,omitempty"`  // default: 0
	Transpositions bool         `json:"transpositions,omitempty"` // default: true
	Rewrite        QueryRewrite `json:"rewrite,omitempty"`
}

func (f *FuzzyQuery) GetQueryName() string { return "fuzzy" }

func (f *FuzzyQuery) MarshalJSON() ([]byte, error) {
	type _FuzzyQuery FuzzyQuery
	data := map[string]*_FuzzyQuery{
		f.Field: (*_FuzzyQuery)(f),
	}
	return json.Marshal(data)
}

var (
	_ Query = (*FuzzyQuery)(nil)
)

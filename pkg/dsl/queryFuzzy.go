package dsl

import "encoding/json"

// FuzzyQuery results a fuzzy query struct
//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-fuzzy-query.html
type Fuzzyquery func(field string, value string, o ...func(all *tFuzzyQuery)) BaseQuery

func (fq *Fuzzyquery) WithFuzziness(fuzziness string) func(all *tFuzzyQuery) {
	return func(all *tFuzzyQuery) {
		all.Fuzziness = fuzziness
	}
}

func (fq *Fuzzyquery) WithMaxExpansion(maxExpansion int) func(all *tFuzzyQuery) {
	return func(all *tFuzzyQuery) {
		all.MaxExpansion = maxExpansion
	}
}

func (fq *Fuzzyquery) WithPrefixLength(prefixLength int) func(all *tFuzzyQuery) {
	return func(all *tFuzzyQuery) {
		all.PrefixLength = prefixLength
	}
}

func (fq *Fuzzyquery) WithTranspositions(transpositions bool) func(all *tFuzzyQuery) {
	return func(all *tFuzzyQuery) {
		all.Transpositions = transpositions
	}
}

func (fq *Fuzzyquery) WithRewrite(rewrite QueryRewrite) func(all *tFuzzyQuery) {
	return func(all *tFuzzyQuery) {
		all.Rewrite = rewrite
	}
}

func NewFuzzyQuery() Fuzzyquery {
	return func(field string, value string, o ...func(all *tFuzzyQuery)) BaseQuery {
		var r = &tFuzzyQuery{
			Field: field,
			Value: value,
		}
		for _, f := range o {
			f(r)
		}
		return r
	}
}

type tFuzzyQuery struct {
	Field          string       `json:"-"`
	Value          string       `json:"value"`
	Fuzziness      string       `json:"fuzziness,omitempty"`
	MaxExpansion   int          `json:"max_expansion,omitempty"`  // default: 50
	PrefixLength   int          `json:"prefix_length,omitempty"`  // default: 0
	Transpositions bool         `json:"transpositions,omitempty"` // default: true
	Rewrite        QueryRewrite `json:"rewrite,omitempty"`
}

func (f *tFuzzyQuery) GetQueryName() string { return "fuzzy" }

func (f *tFuzzyQuery) MarshalJSON() ([]byte, error) {
	type _FuzzyQuery tFuzzyQuery
	data := map[string]*_FuzzyQuery{
		f.Field: (*_FuzzyQuery)(f),
	}
	return json.Marshal(data)
}

var (
	_ BaseQuery = (*tFuzzyQuery)(nil)
)

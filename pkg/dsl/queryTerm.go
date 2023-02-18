package dsl

import "encoding/json"

// TermQuery return a term structure
//
// See complete documentation at https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-term-query.html
type TermQuery func(field string, value string, o ...func(all *tTermQuery)) BaseQuery

func (q *TermQuery) WithBoost(boost float64) func(all *tTermQuery) {
	return func(all *tTermQuery) {
		all.Boost = boost
	}
}

func (q *TermQuery) WithCaseInsensitiv(caseInsensitive bool) func(all *tTermQuery) {
	return func(all *tTermQuery) {
		all.CaseInsensitive = caseInsensitive
	}
}

func NewTermQuery() TermQuery {
	return func(field string, value string, o ...func(all *tTermQuery)) BaseQuery {
		var r = &tTermQuery{
			Field: field,
			Value: value,
		}
		for _, f := range o {
			f(r)
		}
		return r
	}
}

type tTermQuery struct {
	Field           string  `json:"-"`
	Value           string  `json:"value"`
	Boost           float64 `json:"boost,omitempty"`            // default: 1.0
	CaseInsensitive bool    `json:"case_insensitive,omitempty"` // default: false
}

func (*tTermQuery) GetQueryName() string { return "term" }

func (t *tTermQuery) MarshalJSON() ([]byte, error) {
	type _TermQuery tTermQuery
	data := map[string]*_TermQuery{
		t.Field: (*_TermQuery)(t),
	}
	return json.Marshal(data)
}

var (
	_ BaseQuery = (*tTermQuery)(nil)
)

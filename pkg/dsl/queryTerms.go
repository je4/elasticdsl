package dsl

import "encoding/json"

// TermsQuery return a terms structure
//
// See complete documentation at https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-terms-query.html
type TermsQuery func(field string, values []string, o ...func(all *tTermsQuery)) BaseQuery

func (q *TermsQuery) WithBoost(boost float64) func(all *tTermsQuery) {
	return func(all *tTermsQuery) {
		all.Boost = boost
	}
}

func (q *TermsQuery) WithCaseInsensitiv(caseInsensitive bool) func(all *tTermsQuery) {
	return func(all *tTermsQuery) {
		all.CaseInsensitive = caseInsensitive
	}
}

func NewTermsQuery() TermsQuery {
	return func(field string, values []string, o ...func(all *tTermsQuery)) BaseQuery {
		var r = &tTermsQuery{
			Field:  field,
			Values: values,
		}
		for _, f := range o {
			f(r)
		}
		return r
	}
}

type tTermsQuery struct {
	Field           string   `json:"-"`
	Values          []string `json:"-"`
	Boost           float64  `json:"boost,omitempty"`            // default: 1.0
	CaseInsensitive bool     `json:"case_insensitive,omitempty"` // default: false
}

func (*tTermsQuery) GetQueryName() string { return "terms" }

func (t *tTermsQuery) MarshalJSON() ([]byte, error) {
	type _TermsQuery tTermsQuery
	data := map[string][]any{
		t.Field: t.Values,
	}
	if t.Boost != 0.0 {
		data["boost"] = t.Boost
	}
	if t.CaseInsensitive {
		data["case_insensitive"] = t.CaseInsensitive
	}
	return json.Marshal(data)
}

var (
	_ BaseQuery = (*tTermsQuery)(nil)
)

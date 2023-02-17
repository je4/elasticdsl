package dsl

import "encoding/json"

type TermQuery struct {
	Field           string  `json:"-"`
	Value           string  `json:"value"`
	Boost           float64 `json:"boost,omitempty"`            // default: 1.0
	CaseInsensitive bool    `json:"case_insensitive,omitempty"` // default: false
}

func (*TermQuery) GetQueryName() string { return "term" }

func (t *TermQuery) MarshalJSON() ([]byte, error) {
	type _TermQuery TermQuery
	data := map[string]*_TermQuery{
		t.Field: (*_TermQuery)(t),
	}
	return json.Marshal(data)
}

var (
	_ Query = (*TermQuery)(nil)
)

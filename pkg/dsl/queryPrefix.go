package dsl

import "encoding/json"

// https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-prefix-query.html

type PrefixQuery struct {
	Field           string       `json:"-"`
	Value           string       `json:"value"`
	Rewrite         QueryRewrite `json:"rewrite,omitempty"`
	CaseInsensitive bool         `json:"case_insensitive,omitempty"` // default: false
}

func (f *PrefixQuery) GetQueryName() string { return "prefix" }

func (f *PrefixQuery) MarshalJSON() ([]byte, error) {
	type _PrefixQuery PrefixQuery
	data := map[string]*_PrefixQuery{
		f.Field: (*_PrefixQuery)(f),
	}
	return json.Marshal(data)
}

var (
	_ BaseQuery = (*PrefixQuery)(nil)
)

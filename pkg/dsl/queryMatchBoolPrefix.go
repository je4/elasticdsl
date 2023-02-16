package dsl

import (
	"encoding/json"
)

type tMatchBoolPrefixQuery struct {
	Field    string `json:"-"`
	Query    string `json:"query"`
	Analyzer string `json:"analyzer,omitempty"`
}

func (mb *tMatchBoolPrefixQuery) GetQueryName() string { return "match_bool_prefix" }

func (mb *tMatchBoolPrefixQuery) MarshalJSON() ([]byte, error) {
	type _tMatchBoolPrefixQuery tMatchBoolPrefixQuery
	data := map[string]*_tMatchBoolPrefixQuery{
		mb.Field: (*_tMatchBoolPrefixQuery)(mb),
	}
	return json.Marshal(data)
}

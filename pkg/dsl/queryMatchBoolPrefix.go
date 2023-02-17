package dsl

import (
	"encoding/json"
)

type MatchBoolPrefixQuery struct {
	Field    string `json:"-"`
	Query    string `json:"query"`
	Analyzer string `json:"analyzer,omitempty"`
}

func (mb *MatchBoolPrefixQuery) GetQueryName() string { return "match_bool_prefix" }

func (mb *MatchBoolPrefixQuery) MarshalJSON() ([]byte, error) {
	type _tMatchBoolPrefixQuery MatchBoolPrefixQuery
	data := map[string]*_tMatchBoolPrefixQuery{
		mb.Field: (*_tMatchBoolPrefixQuery)(mb),
	}
	return json.Marshal(data)
}

var (
	_ Query = (*MatchBoolPrefixQuery)(nil)
)

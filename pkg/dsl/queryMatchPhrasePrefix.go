package dsl

import "encoding/json"

type tMatchPhrasePrefixQuery struct {
	Field          string          `json:"-"`
	Query          string          `json:"query"`
	Analyzer       string          `json:"analyzer,omitempty"`
	MaxExpansions  int             `json:"max_expansions,omitempty"`   // default: 50
	Slop           int             `json:"slop,omitempty"`             // default: 0
	ZeroTermsQuery tZeroTermsQuery `json:"zero_terms_query,omitempty"` // default: ZeroTermsQueryNONE
}

func (mpp *tMatchPhrasePrefixQuery) GetQueryName() string { return "match_phrase_prefix" }

func (mpp *tMatchPhrasePrefixQuery) MarshalJSON() ([]byte, error) {
	type _tMatchPhrasePrefixQuery tMatchPhrasePrefixQuery
	data := map[string]*_tMatchPhrasePrefixQuery{
		mpp.Field: (*_tMatchPhrasePrefixQuery)(mpp),
	}
	return json.Marshal(data)
}

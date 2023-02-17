package dsl

import "encoding/json"

type MatchPhrasePrefixQuery struct {
	Field          string              `json:"-"`
	Query          string              `json:"query"`
	Analyzer       string              `json:"analyzer,omitempty"`
	MaxExpansions  int                 `json:"max_expansions,omitempty"`   // default: 50
	Slop           int                 `json:"slop,omitempty"`             // default: 0
	ZeroTermsQuery QueryZeroTermsQuery `json:"zero_terms_query,omitempty"` // default: QeruyZeroTermsQueryNONE
}

func (mpp *MatchPhrasePrefixQuery) GetQueryName() string { return "match_phrase_prefix" }

func (mpp *MatchPhrasePrefixQuery) MarshalJSON() ([]byte, error) {
	type _tMatchPhrasePrefixQuery MatchPhrasePrefixQuery
	data := map[string]*_tMatchPhrasePrefixQuery{
		mpp.Field: (*_tMatchPhrasePrefixQuery)(mpp),
	}
	return json.Marshal(data)
}

var (
	_ Query = (*MatchPhrasePrefixQuery)(nil)
)

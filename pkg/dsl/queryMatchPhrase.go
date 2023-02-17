package dsl

import "encoding/json"

type MatchPhraseQuery struct {
	Field    string `json:"-"`
	Query    string `json:"query"`
	Analyzer string `json:"analyzer,omitempty"`
}

func (mp *MatchPhraseQuery) MarshalJSON() ([]byte, error) {
	type _tMatchBoolPrefixQuery MatchBoolPrefixQuery
	data := map[string]*_tMatchBoolPrefixQuery{
		mp.Field: (*_tMatchBoolPrefixQuery)(mp),
	}
	return json.Marshal(data)
}

func (mp *MatchPhraseQuery) GetQueryName() string { return "match_phrase" }

var (
	_ Query = (*MatchPhraseQuery)(nil)
)

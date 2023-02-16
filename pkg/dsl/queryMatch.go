package dsl

import "encoding/json"

// https://www.elastic.co/guide/en/elasticsearch/reference/8.6/query-dsl-match-query.html

type tMatchQuery struct {
	Field                           string          `json:"-"`
	Query                           string          `json:"query"`
	Analyzer                        string          `json:"analyzer,omitempty"`
	AutoGenerateSynonymsPhraseQuery bool            `json:"auto_generate_synonyms_phrase_query,omitempty"`
	Fuzziness                       string          `json:"fuzziness,omitempty"`
	MaxExpansions                   int             `json:"max_expansions,omitempty"`       // default: 50
	PrefixLength                    int             `json:"prefix_length,omitempty"`        // default: 0
	FuzzyTranspositions             bool            `json:"fuzzy_transpositions,omitempty"` // default: true
	FuzzyRewrite                    string          `json:"fuzzy_rewrite,omitempty"`
	Lenient                         bool            `json:"lenient,omitempty"`  // default: false
	Operator                        tOperator       `json:"operator,omitempty"` // default: OperatorOR
	MinimumShouldMatch              string          `json:"minimum_should_match,omitempty"`
	ZeroTermsQuery                  tZeroTermsQuery `json:"zero_terms_query,omitempty"` // default: ZeroTermsQueryNONE
}

func (m *tMatchQuery) GetQueryName() string { return "match" }

func (m *tMatchQuery) MarshalJSON() ([]byte, error) {
	type _tMatchQuery tMatchQuery
	data := map[string]*_tMatchQuery{
		m.Field: (*_tMatchQuery)(m),
	}
	return json.Marshal(data)
}

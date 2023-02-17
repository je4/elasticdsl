package dsl

import (
	"encoding/json"
	"strings"
)

// https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-simple-query-string-query.html

type SimpleQueryStringQueryFlags []QuerySimpleQueryStringFlag

func (flags SimpleQueryStringQueryFlags) MarshalJSON() ([]byte, error) {
	var strFlags = []string{}
	for _, f := range flags {
		strFlags = append(strFlags, string(f))
	}
	data := strings.Join(strFlags, "|")
	return json.Marshal(data)
}

type SimpleQueryStringQuery struct {
	Query                           string                      `json:"query"`
	Fields                          []string                    `json:"fields,omitempty"`
	DefaultOperator                 QueryOperator               `json:"default_operator,omitempty"` // default: QueryOperatorOR
	AnalyzeWildcard                 bool                        `json:"analyze_wildcard,omitempty"` // default: false
	Analyzer                        string                      `json:"analyzer,omitempty"`
	AutoGenerateSynonymsPhraseQuery bool                        `json:"auto_generate_synonyms_phrase_query,omitempty"`
	Flags                           SimpleQueryStringQueryFlags `json:"flags,omitempty"`
	FuzzyMaxExpansions              int                         `json:"fuzzy_max_expansions,omitempty"` // default: 50
	FuzzyPrefixLength               int                         `json:"fuzzy_prefix_length,omitempty"`  // default: 0
	FuzzyTranspositions             bool                        `json:"fuzzy_transpositions,omitempty"` // default: true
	Lenient                         bool                        `json:"lenient,omitempty"`              // default: false
	MinimumShouldMatch              string                      `json:"minimum_should_match,omitempty"`
	QuoteFieldSuffix                string                      `json:"quote_field_suffix,omitempty"`
}

func (*SimpleQueryStringQuery) GetQueryName() string { return "simple_query_string" }

var (
	_ Query = (*SimpleQueryStringQuery)(nil)
)

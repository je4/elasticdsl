package dsl

// https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-multi-match-query.html

type MultiMatchQuery struct {
	Fields                          []string            `json:"fields,omitempty"`
	Query                           string              `json:"query"`
	Type                            QueryMultiMatchType `json:"type,omitempty"` // default: QueryMultiMatchTypeBestFields
	Analyzer                        string              `json:"analyzer,omitempty"`
	AutoGenerateSynonymsPhraseQuery bool                `json:"auto_generate_synonyms_phrase_query,omitempty"`
	Fuzziness                       string              `json:"fuzziness,omitempty"`
	MaxExpansions                   int                 `json:"max_expansions,omitempty"`       // default: 50
	PrefixLength                    int                 `json:"prefix_length,omitempty"`        // default: 0
	FuzzyTranspositions             bool                `json:"fuzzy_transpositions,omitempty"` // default: true
	FuzzyRewrite                    string              `json:"fuzzy_rewrite,omitempty"`
	Lenient                         bool                `json:"lenient,omitempty"`  // default: false
	Operator                        QueryOperator       `json:"operator,omitempty"` // default: QueryOperatorOR
	MinimumShouldMatch              string              `json:"minimum_should_match,omitempty"`
	ZeroTermsQuery                  QueryZeroTermsQuery `json:"zero_terms_query,omitempty"` // default: QeruyZeroTermsQueryNONE
	TieBreaker                      float64             `json:"tie_breaker,omitempty"`      // default: 0.0
}

func (*MultiMatchQuery) GetQueryName() string { return "multi_match" }

var (
	_ Query = (*MultiMatchQuery)(nil)
)

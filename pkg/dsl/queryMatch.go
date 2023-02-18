package dsl

import "encoding/json"

// MatchQuery results a match query struct
//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/8.6/query-dsl-match-query.html
type MatchQuery func(field string, query string, o ...func(all *tMatchQuery)) BaseQuery

func (q *MatchQuery) WithAnalyzer(v string) func(all *tMatchQuery) {
	return func(all *tMatchQuery) {
		all.Analyzer = v
	}
}

func (q *MatchQuery) WithAutoGenerateSynonymsPhraseQuery(v bool) func(all *tMatchQuery) {
	return func(all *tMatchQuery) {
		all.AutoGenerateSynonymsPhraseQuery = v
	}
}

func (q *MatchQuery) WithFuzziness(v string) func(all *tMatchQuery) {
	return func(all *tMatchQuery) {
		all.Fuzziness = v
	}
}

func (q *MatchQuery) WithMaxExpansions(v int) func(all *tMatchQuery) {
	return func(all *tMatchQuery) {
		all.MaxExpansions = v
	}
}

func (q *MatchQuery) WithPrefixLength(v int) func(all *tMatchQuery) {
	return func(all *tMatchQuery) {
		all.PrefixLength = v
	}
}

func (q *MatchQuery) WithFuzzyTranspositions(v bool) func(all *tMatchQuery) {
	return func(all *tMatchQuery) {
		all.FuzzyTranspositions = v
	}
}

func (q *MatchQuery) WithFuzzyRewrite(v string) func(all *tMatchQuery) {
	return func(all *tMatchQuery) {
		all.FuzzyRewrite = v
	}
}

func (q *MatchQuery) WithLenient(v bool) func(all *tMatchQuery) {
	return func(all *tMatchQuery) {
		all.Lenient = v
	}
}

func (q *MatchQuery) WithOperator(v QueryOperator) func(all *tMatchQuery) {
	return func(all *tMatchQuery) {
		all.Operator = v
	}
}

func (q *MatchQuery) WithMinimumShouldMatch(v *MinimumShouldMatch) func(all *tMatchQuery) {
	return func(all *tMatchQuery) {
		all.MinimumShouldMatch = v
	}
}

func (q *MatchQuery) WithZeroTermsQuery(v QueryZeroTermsQuery) func(all *tMatchQuery) {
	return func(all *tMatchQuery) {
		all.ZeroTermsQuery = v
	}
}

func NewMatchQuery() MatchQuery {
	return func(field string, query string, o ...func(all *tMatchQuery)) BaseQuery {
		var r = &tMatchQuery{
			Field: field,
			Query: query,
		}
		for _, f := range o {
			f(r)
		}
		return r
	}
}

type tMatchQuery struct {
	Field                           string              `json:"-"`
	Query                           string              `json:"query"`
	Analyzer                        string              `json:"analyzer,omitempty"`
	AutoGenerateSynonymsPhraseQuery bool                `json:"auto_generate_synonyms_phrase_query,omitempty"`
	Fuzziness                       string              `json:"fuzziness,omitempty"`
	MaxExpansions                   int                 `json:"max_expansions,omitempty"`       // default: 50
	PrefixLength                    int                 `json:"prefix_length,omitempty"`        // default: 0
	FuzzyTranspositions             bool                `json:"fuzzy_transpositions,omitempty"` // default: true
	FuzzyRewrite                    string              `json:"fuzzy_rewrite,omitempty"`
	Lenient                         bool                `json:"lenient,omitempty"`  // default: false
	Operator                        QueryOperator       `json:"operator,omitempty"` // default: QueryOperatorOR
	MinimumShouldMatch              *MinimumShouldMatch `json:"minimum_should_match,omitempty"`
	ZeroTermsQuery                  QueryZeroTermsQuery `json:"zero_terms_query,omitempty"` // default: QeruyZeroTermsQueryNONE
}

func (m *tMatchQuery) GetQueryName() string { return "match" }

func (m *tMatchQuery) MarshalJSON() ([]byte, error) {
	type _tMatchQuery tMatchQuery
	data := map[string]*_tMatchQuery{
		m.Field: (*_tMatchQuery)(m),
	}
	return json.Marshal(map[string]any{
		m.GetQueryName(): data,
	})
}

var (
	_ BaseQuery = (*tMatchQuery)(nil)
)

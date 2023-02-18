package dsl

// SimpleQueryStringQuery results a simple_query_string struct
//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-simple-query-string-query.html
type SimpleQueryStringQuery func(query string, o ...func(all *tSimpleQueryStringQuery)) BaseQuery

func (sqsq *SimpleQueryStringQuery) WithFields(fields ...string) func(all *tSimpleQueryStringQuery) {
	return func(all *tSimpleQueryStringQuery) {
		all.Fields = fields
	}
}

func (sqsq *SimpleQueryStringQuery) WithDefaultOperator(defaultOperator QueryOperator) func(all *tSimpleQueryStringQuery) {
	return func(all *tSimpleQueryStringQuery) {
		all.DefaultOperator = defaultOperator
	}
}

func (sqsq *SimpleQueryStringQuery) WithAnalyzeWildcard(analyzeWildcard bool) func(all *tSimpleQueryStringQuery) {
	return func(all *tSimpleQueryStringQuery) {
		all.AnalyzeWildcard = analyzeWildcard
	}
}

func (sqsq *SimpleQueryStringQuery) WithAnalyzer(analyzer string) func(all *tSimpleQueryStringQuery) {
	return func(all *tSimpleQueryStringQuery) {
		all.Analyzer = analyzer
	}
}

func (sqsq *SimpleQueryStringQuery) WithAutoGenerateSynonymsPhraseQuery(autoGenerateSynonymsPhraseQuery bool) func(all *tSimpleQueryStringQuery) {
	return func(all *tSimpleQueryStringQuery) {
		all.AutoGenerateSynonymsPhraseQuery = autoGenerateSynonymsPhraseQuery
	}
}

func (sqsq *SimpleQueryStringQuery) WithFlags(flags ...QuerySimpleQueryStringFlag) func(all *tSimpleQueryStringQuery) {
	return func(all *tSimpleQueryStringQuery) {
		all.Flags = flags
	}
}

func (sqsq *SimpleQueryStringQuery) WithFuzzyMaxExpansions(FuzzyMaxExpansions int) func(all *tSimpleQueryStringQuery) {
	return func(all *tSimpleQueryStringQuery) {
		all.FuzzyMaxExpansions = FuzzyMaxExpansions
	}
}

func (sqsq *SimpleQueryStringQuery) WithFuzzyPrefixLength(FuzzyPrefixLength int) func(all *tSimpleQueryStringQuery) {
	return func(all *tSimpleQueryStringQuery) {
		all.FuzzyPrefixLength = FuzzyPrefixLength
	}
}

func (sqsq *SimpleQueryStringQuery) WithFuzzyTranspositions(FuzzyTranspositions bool) func(all *tSimpleQueryStringQuery) {
	return func(all *tSimpleQueryStringQuery) {
		all.FuzzyTranspositions = FuzzyTranspositions
	}
}

func (sqsq *SimpleQueryStringQuery) WithLenient(Lenient bool) func(all *tSimpleQueryStringQuery) {
	return func(all *tSimpleQueryStringQuery) {
		all.Lenient = Lenient
	}
}

func (sqsq *SimpleQueryStringQuery) WithMinimumShouldMatch(MinimumShouldMatch *MinimumShouldMatch) func(all *tSimpleQueryStringQuery) {
	return func(all *tSimpleQueryStringQuery) {
		all.MinimumShouldMatch = MinimumShouldMatch
	}
}

func (sqsq *SimpleQueryStringQuery) WithQuoteFieldSuffix(QuoteFieldSuffix string) func(all *tSimpleQueryStringQuery) {
	return func(all *tSimpleQueryStringQuery) {
		all.QuoteFieldSuffix = QuoteFieldSuffix
	}
}

func NewSimpleQueryStringQuery() SimpleQueryStringQuery {
	return func(query string, o ...func(all *tSimpleQueryStringQuery)) BaseQuery {
		var r = &tSimpleQueryStringQuery{
			Query: query,
		}
		for _, f := range o {
			f(r)
		}
		return r
	}
}

type tSimpleQueryStringQuery struct {
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
	MinimumShouldMatch              *MinimumShouldMatch         `json:"minimum_should_match,omitempty"`
	QuoteFieldSuffix                string                      `json:"quote_field_suffix,omitempty"`
}

func (*tSimpleQueryStringQuery) GetQueryName() string { return "simple_query_string" }

var (
	_ BaseQuery = (*tSimpleQueryStringQuery)(nil)
)

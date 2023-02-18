package dsl

// https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-query-string-query.html#query-dsl-query-string-query

type StringQuery struct {
	Query                           string        `json:"query"`
	DefaultField                    string        `json:"default_field,omitempty"`
	AllowLeadingWildcard            bool          `json:"allow_leading_wildcard,omitempty"` // default: true
	AnalyzeWildcard                 bool          `json:"analyze_wildcard,omitempty"`       // default: false
	Analyzer                        string        `json:"analyzer,omitempty"`
	AutoGenerateSynonymsPhraseQuery bool          `json:"auto_generate_synonyms_phrase_query,omitempty"`
	Boost                           float64       `json:"boost,omitempty"`                      // default: 1.0
	DefaultOperator                 QueryOperator `json:"default_operator,omitempty"`           // default: QueryOperatorOR
	EnablePositionIncrements        bool          `json:"enable_position_increments,omitempty"` // default: true
	Fields                          []string      `json:"fields,omitempty"`
	Fuzziness                       string        `json:"fuzziness,omitempty"`
	FuzzyMaxExpansions              int           `json:"fuzzy_max_expansions,omitempty"`    // default: 50
	FuzzyPrefixLength               int           `json:"fuzzy_prefix_length,omitempty"`     // default: 0
	FuzzyTranspositions             bool          `json:"fuzzy_transpositions,omitempty"`    // default: true
	Lenient                         bool          `json:"lenient,omitempty"`                 // default: false
	MaxDeterminizedStates           int           `json:"max_determinized_states,omitempty"` // default: 10000
	MinimumShouldMatch              string        `json:"minimum_should_match,omitempty"`
	QuoteAnalyzer                   string        `json:"quote_analyzer,omitempty"`
	PhraseSlop                      int           `json:"phrase_slop,omitempty"` // default: 0
	QuoteFieldSuffix                string        `json:"quote_field_suffix,omitempty"`
	Rewrite                         string        `json:"rewrite,omitempty"`
	TimeZone                        string        `json:"time_zone,omitempty"`
}

func (*StringQuery) GetQueryName() string { return "query_string" }

var (
	_ BaseQuery = (*StringQuery)(nil)
)

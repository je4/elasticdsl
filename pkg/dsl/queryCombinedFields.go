package dsl

type CombinedFieldsQuery struct {
	Fields                          []string            `json:"fields"`
	Query                           string              `json:"query"`
	AutoGenerateSynonymsPhraseQuery bool                `json:"auto_generate_synonyms_phrase_query,omitempty"`
	Operator                        QueryOperator       `json:"operator,omitempty"` // default: QueryOperatorOR
	MinimumShouldMatch              string              `json:"minimum_should_match,omitempty"`
	ZeroTermsQuery                  QueryZeroTermsQuery `json:"zero_terms_query,omitempty"` // default: QeruyZeroTermsQueryNONE
}

func (*CombinedFieldsQuery) GetQueryName() string { return "combined_fields" }

var (
	_ Query = (*CombinedFieldsQuery)(nil)
)

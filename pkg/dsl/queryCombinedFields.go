package dsl

import "encoding/json"

// CombindedFieldQuery results a combined_field query struct
//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-combined-fields-query.html
type CombinedFieldQuery func(fields []string, query string, o ...func(all *tCombinedFieldsQuery)) BaseQuery

func (q *CombinedFieldQuery) WithAutoGenerateSynonymsPhraseQuery(b bool) func(all *tCombinedFieldsQuery) {
	return func(all *tCombinedFieldsQuery) {
		all.AutoGenerateSynonymsPhraseQuery = b
	}
}

func (q *CombinedFieldQuery) WithOperator(o QueryOperator) func(all *tCombinedFieldsQuery) {
	return func(all *tCombinedFieldsQuery) {
		all.Operator = o
	}
}

func (q *CombinedFieldQuery) WithMinimumShouldMatch(v string) func(all *tCombinedFieldsQuery) {
	return func(all *tCombinedFieldsQuery) {
		all.MinimumShouldMatch = v
	}
}

func (q *CombinedFieldQuery) WithZeroTermsQuery(v QueryZeroTermsQuery) func(all *tCombinedFieldsQuery) {
	return func(all *tCombinedFieldsQuery) {
		all.ZeroTermsQuery = v
	}
}

func NewCombinedFieldQuery() CombinedFieldQuery {
	return func(fields []string, query string, o ...func(all *tCombinedFieldsQuery)) BaseQuery {
		var r = &tCombinedFieldsQuery{
			Fields: fields,
			Query:  query,
		}
		for _, f := range o {
			f(r)
		}
		return r
	}
}

type tCombinedFieldsQuery struct {
	Fields                          []string            `json:"fields"`
	Query                           string              `json:"query"`
	AutoGenerateSynonymsPhraseQuery bool                `json:"auto_generate_synonyms_phrase_query,omitempty"`
	Operator                        QueryOperator       `json:"operator,omitempty"` // default: QueryOperatorOR
	MinimumShouldMatch              string              `json:"minimum_should_match,omitempty"`
	ZeroTermsQuery                  QueryZeroTermsQuery `json:"zero_terms_query,omitempty"` // default: QeruyZeroTermsQueryNONE
}

func (*tCombinedFieldsQuery) GetQueryName() string { return "combined_fields" }

func (q *tCombinedFieldsQuery) MarshalJSON() ([]byte, error) {
	type _tCombinedFieldsQuery tCombinedFieldsQuery
	return json.Marshal(map[string]*_tCombinedFieldsQuery{
		q.GetQueryName(): (*_tCombinedFieldsQuery)(q),
	})
}

var (
	_ BaseQuery = (*tCombinedFieldsQuery)(nil)
)

package dsl

type API struct {
	Query                  Query
	MatchAllQuery          MatchAllQuery
	CombinedFieldQuery     CombinedFieldQuery
	ExistsQuery            ExistsQuery
	BoolQuery              BoolQuery
	MatchQuery             MatchQuery
	MatchNoneQuery         MatchNoneQuery
	Fuzzyquery             Fuzzyquery
	SimpleQueryStringQuery SimpleQueryStringQuery
	NestedQuery            NestedQuery
	TermQuery              TermQuery
	TermsQuery             TermsQuery
	Search                 Search
	Aggs                   Aggs
	AggNested              AggNested
	AggTerms               AggTerms
	AggFilter              AggFilter
	AggScript              AggScript
	IndicesBoost           IndicesBoost
	Fields                 Fields
	FieldsField            FieldsField
}

func NewApi() *API {
	api := &API{
		Query:                  NewQuery(),
		MatchAllQuery:          NewMatchAllQuery(),
		CombinedFieldQuery:     NewCombinedFieldQuery(),
		ExistsQuery:            NewExistsQuery(),
		BoolQuery:              NewBoolQuery(),
		MatchQuery:             NewMatchQuery(),
		MatchNoneQuery:         NewMatchNoneQuery(),
		Fuzzyquery:             NewFuzzyQuery(),
		SimpleQueryStringQuery: NewSimpleQueryStringQuery(),
		NestedQuery:            NewNestedQuery(),
		TermQuery:              NewTermQuery(),
		TermsQuery:             NewTermsQuery(),
		Search:                 NewSearch(),
		Aggs:                   NewAggs(),
		AggNested:              NewAggNested(),
		AggTerms:               NewAggTerms(),
		AggFilter:              NewAggFilter(),
		AggScript:              NewAggScript(),
		IndicesBoost:           NewIndicesBoost(),
		Fields:                 NewFields(),
		FieldsField:            NewFieldsField(),
	}
	return api
}

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
	IndicesBoost           IndicesBoost
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
		IndicesBoost:           NewIndicesBoost(),
	}
	return api
}

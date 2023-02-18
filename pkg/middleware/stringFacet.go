package middleware

import "github.com/je4/elasticdsl/v2/pkg/dsl"

type StringFacet struct {
	Name   string
	Values []string
}

func (sf *StringFacet) GetName() string { return sf.Name }

func (sf *StringFacet) GetAggs(api *dsl.API) any {
	var aggs = api.Aggs(
		api.Aggs.WithQueries(
			api.BoolQuery(api.BoolQuery.WithMust(
				api.TermQuery("name", sf.Name),
				api.TermsQuery("string", sf.Values),
			),
			),
		)
	)
	return aggs
}

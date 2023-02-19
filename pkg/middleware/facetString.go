package middleware

import "github.com/je4/elasticdsl/v2/pkg/dsl"

type StringFacet struct {
	Name   string
	Values []string
}

func (sf *StringFacet) GetName() string { return sf.Name }

func (sf *StringFacet) GetAgg(api *dsl.API) dsl.Aggregation {
	var aggs = api.AggNested(
		"facet.strings",
		api.Aggs(
			api.Aggs.WithAggregations(
				api.AggTerms(
					"facet_strings_"+sf.Name,
					"facet.strings."+sf.Name,
				),
			),
		),
	)
	return aggs
}

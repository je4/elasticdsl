package middleware

import "github.com/je4/elasticdsl/v2/pkg/dsl"

type StringFacet struct {
	Name   string
	Values []string
}

func (sf *StringFacet) GetName() string { return sf.Name }

func (sf *StringFacet) GetAgg(api *dsl.API) dsl.BaseAgg {
	var agg = api.AggFilter(
		"facet-strings-"+sf.Name+"-filter",
		"facet.strings.name",
		sf.Name,
		api.AggFilter.WithAggs(
			api.AggTerms(
				"facet-strings-"+sf.Name+"-terms",
				"facet.strings.string",
			),
		),
	)
	return agg
}

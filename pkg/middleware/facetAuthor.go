package middleware

import "github.com/je4/elasticdsl/v2/pkg/dsl"

type AuthorFacet struct {
	Name        string
	Identifiers []string
}

func (af *AuthorFacet) GetName() string { return af.Name }

func (af *AuthorFacet) GetAgg(api *dsl.API) dsl.BaseAgg {
	var aggs = api.AggNested(
		"facet_object_"+af.Name,
		"facet.objects",
		api.Aggs(
			api.Aggs.WithAggregations(
				api.AggTerms(
					"facet_objects_terms_"+af.Name,
					"facet.objects.",
				),
			),
		),
	)
	return aggs
}

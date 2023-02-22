package middleware

import (
	"emperror.dev/errors"
	"github.com/je4/elasticdsl/v2/pkg/dsl"
	"github.com/je4/elasticdsl/v2/pkg/elastic"
)

type Middleware struct {
	es elastic.Client
}

type SearchFacets struct {
	StringFacets StringFacets
	PersonFacets PersonFacets
}

func (m *Middleware) Search(index string, facets *SearchFacets) (*Result, error) {
	var api = m.es.GetDSLAPI()

	var aggs = []dsl.BaseAgg{}
	if facets.StringFacets != nil {
		aggs = append(aggs, facets.StringFacets.GetAgg(api))
	}
	if facets.PersonFacets != nil {
		aggs = append(aggs, facets.PersonFacets.GetAgg(api))
	}

	search := api.Search(
		api.Search.WithQuery(
			api.Query(
				api.MatchAllQuery(
					api.MatchAllQuery.WithBoost(0.5),
				),
			),
		),
		api.Search.WithAggs(aggs...),
		/*
			api.Search.WithFields(
				api.Fields(
					api.Fields.WithFields(
						api.FieldsField("mapping.*"),
						api.FieldsField("acl.*"),
						api.FieldsField("sets"),
						api.FieldsField("flags"),
					),
				),
			),
		*/
	)

	result, err := m.es.Search(index, search)
	if err != nil {
		return nil, errors.Wrap(err, "cannot query")
	}
	strf, err := facets.StringFacets.UnmarshalJSON(result.Aggregations["facet-string"])
	if err != nil {
		return nil, errors.Wrapf(err, "cannot unmarshal '%s'", string(result.Aggregations["facet-string"]))
	}
	var _ = strf
	pf, err := facets.StringFacets.UnmarshalJSON(result.Aggregations["facet-object"])
	if err != nil {
		return nil, errors.Wrapf(err, "cannot unmarshal '%s'", string(result.Aggregations["facet-person"]))
	}
	var _ = pf
	var ret = &Result{}
	return ret, nil
}

func NewMiddleware(es elastic.Client) *Middleware {
	var m = &Middleware{es: es}
	return m
}

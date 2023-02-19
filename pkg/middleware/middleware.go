package middleware

import (
	"emperror.dev/errors"
	"github.com/je4/elasticdsl/v2/pkg/dsl"
	"github.com/je4/elasticdsl/v2/pkg/elastic"
)

type Middleware struct {
	es elastic.Client
}

func (m *Middleware) Search(index string, facets []Facet) (*Result, error) {
	var api = m.es.GetDSLAPI()

	var aggs = []dsl.Aggregation{}
	for _, facet := range facets {
		aggs = append(aggs, facet.GetAgg(api))
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
	)

	result, err := m.es.Search(index, search)
	if err != nil {
		return nil, errors.Wrap(err, "cannot query")
	}
	_ = result
	var ret = &Result{}
	return ret, nil
}

func NewMiddleware(es elastic.Client) *Middleware {
	var m = &Middleware{es: es}
	return m
}

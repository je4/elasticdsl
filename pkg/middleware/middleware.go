package middleware

import (
	"emperror.dev/errors"
	"github.com/je4/elasticdsl/v2/pkg/dsl"
	"github.com/je4/elasticdsl/v2/pkg/elastic"
)

type Middleware struct {
	es elastic.Client
}

func (m *Middleware) Search(index string) (*Result, error) {
	api := m.es.GetDSLAPI()
	search := api.Search(
		api.Search.WithQuery(
			api.Query(
				api.BoolQuery(
					api.BoolQuery.WithMust(
						api.MatchAllQuery(
							api.MatchAllQuery.WithBoost(0.5),
						),
					),
					api.BoolQuery.WithMinimumShouldMatch(
						&dsl.MinimumShouldMatch{IntValue: 1},
					),
				),
			),
		),
		api.Search.WithAggs(
			api.Aggs(),
		),
		api.Search.WithIndicesBoost(
			api.IndicesBoost(
				api.IndicesBoost.AppendIndex(index, 2.1),
			),
		),
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

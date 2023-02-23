package middleware

import (
	"emperror.dev/errors"
	"encoding/json"
	"fmt"
	"github.com/je4/elasticdsl/v2/pkg/dsl"
	"github.com/je4/elasticdsl/v2/pkg/elastic"
)

type PersonFacets []*PersonFacet

type PersonFacetsResult struct {
	DocCount int `json:"-"`
	Facets   map[string][]*PersonFacetResult
}

func (PersonFacets) GetName() string {
	return "facet-person"
}

func (sf PersonFacets) GetAgg(api *dsl.API) dsl.BaseAgg {
	var facets = []dsl.BaseAgg{}
	for _, facet := range sf {
		facets = append(facets, facet.GetAgg(api))
	}
	var aggs = api.AggNested(
		sf.GetName(),
		"facet.objects",
		api.Aggs(
			api.Aggs.WithAggregations(
				facets...,
			),
		),
	)
	return aggs

}

func (sf PersonFacets) FromJSON(dataBytes []byte) (*PersonFacetsResult, error) {
	var data = map[string]elastic.JSONDummy{}
	if err := json.Unmarshal(dataBytes, &data); err != nil {
		return nil, errors.Wrapf(err, "cannot unmarshal '%s'", string(dataBytes))
	}
	var sfr = &PersonFacetsResult{
		Facets: map[string][]*PersonFacetResult{},
	}
	if docCountBytes, ok := data["doc_count"]; ok {
		if err := json.Unmarshal(docCountBytes, &sfr.DocCount); err != nil {
			return nil, errors.Wrapf(err, "cannot unmarshal '%s'", string(docCountBytes))
		}
	}
	for _, facet := range sf {
		key := fmt.Sprintf("facet-persons-%s-filter", facet.GetName())
		if vals, ok := data[key]; ok {
			r, err := facet.FromJSON(vals)
			if err != nil {
				return nil, errors.Wrapf(err, "cannot unmarshal '%s'", string(vals))
			}
			sfr.Facets[facet.GetName()] = r
		}
	}
	return sfr, nil
}

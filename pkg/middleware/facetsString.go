package middleware

import (
	"emperror.dev/errors"
	"encoding/json"
	"fmt"
	"github.com/je4/elasticdsl/v2/pkg/dsl"
	"github.com/je4/elasticdsl/v2/pkg/elastic"
)

type StringFacets []*StringFacet

type StringFacetsResult struct {
	DocCount int `json:"-"`
	Facets   map[string][]*StringFacetResult
}

func (StringFacets) GetName() string {
	return "facet-string"
}

func (sf StringFacets) GetAgg(api *dsl.API) dsl.BaseAgg {
	var facets = []dsl.BaseAgg{}
	for _, facet := range sf {
		facets = append(facets, facet.GetAgg(api))
	}
	var aggs = api.AggNested(
		sf.GetName(),
		"facet.strings",
		api.Aggs(
			api.Aggs.WithAggregations(
				facets...,
			),
		),
	)
	return aggs

}

func (sf StringFacets) UnmarshalJSON(dataBytes []byte) (*StringFacetsResult, error) {
	var data = map[string]elastic.JSONDummy{}
	if err := json.Unmarshal(dataBytes, &data); err != nil {
		return nil, errors.Wrapf(err, "cannot unmarshal '%s'", string(dataBytes))
	}
	var sfr = &StringFacetsResult{
		Facets: map[string][]*StringFacetResult{},
	}
	if docCountBytes, ok := data["doc_count"]; ok {
		if err := json.Unmarshal(docCountBytes, &sfr.DocCount); err != nil {
			return nil, errors.Wrapf(err, "cannot unmarshal '%s'", string(docCountBytes))
		}
	}
	for _, facet := range sf {
		key := fmt.Sprintf("facet-strings-%s-filter", facet.GetName())
		if vals, ok := data[key]; ok {
			r, err := facet.UnmarshalJSON(vals)
			if err != nil {
				return nil, errors.Wrapf(err, "cannot unmarshal '%s'", string(vals))
			}
			sfr.Facets[facet.GetName()] = r
		}
	}
	return sfr, nil
}

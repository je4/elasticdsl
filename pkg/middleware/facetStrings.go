package middleware

import (
	"emperror.dev/errors"
	"encoding/json"
	"github.com/je4/elasticdsl/v2/pkg/dsl"
	"github.com/je4/elasticdsl/v2/pkg/elastic"
)

type StringFacets []*StringFacet

type StringFacetsResult struct {
	DocCount int `json:"-"`
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

func (sf StringFacets) UnmarshalJSON([]byte) error {
	var data = map[string]elastic.JSONDummy{}
	var sfr = &StringFacetsResult{}
	if docCountBytes, ok := data["doc_count"]; ok {
		if err := json.Unmarshal(docCountBytes, &sfr.DocCount); err != nil {
			return errors.Wrapf(err, "cannot unmarshal '%s'", string(docCountBytes))
		}
	}
	return nil
}

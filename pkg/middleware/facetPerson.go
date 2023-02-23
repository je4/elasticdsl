package middleware

import (
	"emperror.dev/errors"
	"encoding/json"
	"fmt"
	"github.com/je4/elasticdsl/v2/pkg/dsl"
	"github.com/je4/elasticdsl/v2/pkg/elastic"
)

type PersonFacet struct {
	Name   string
	Values []string
}

type PersonFacetResult struct {
	DocCount int    `json:"doc_count"`
	Key      string `json:"key"`
}

func (sf *PersonFacet) GetName() string { return sf.Name }

func (sf *PersonFacet) GetAgg(api *dsl.API) dsl.BaseAgg {
	var agg = api.AggFilter(
		"facet-persons-"+sf.Name+"-filter",
		"facet.objects.name",
		sf.Name,
		api.AggFilter.WithAggs(
			api.AggTerms(
				"facet-objects-"+sf.Name+"-terms",
				"facet.objects.objects",
			),
		),
	)
	return agg
}

func (sf PersonFacet) FromJSON(dataBytes []byte) ([]*PersonFacetResult, error) {
	var data = map[string]elastic.JSONDummy{}
	if err := json.Unmarshal(dataBytes, &data); err != nil {
		return nil, errors.Wrapf(err, "cannot unmarshal '%s'", string(dataBytes))
	}
	var sfr = &PersonFacetsResult{}
	if docCountBytes, ok := data["doc_count"]; ok {
		if err := json.Unmarshal(docCountBytes, &sfr.DocCount); err != nil {
			return nil, errors.Wrapf(err, "cannot unmarshal '%s'", string(docCountBytes))
		}
	}
	key := fmt.Sprintf("facet-persons-%s-terms", sf.GetName())
	if terms, ok := data[key]; ok {
		termsStruct := &struct {
			DocCountErrorUpperBound int                  `json:"doc_count_error_upper_bound,omitempty"`
			SumOtherDocCount        int                  `json:"sum_other_doc_count,omitempty"`
			Buckets                 []*PersonFacetResult `json:"buckets"`
		}{}
		if err := json.Unmarshal(terms, termsStruct); err != nil {
			return nil, errors.Wrapf(err, "cannot unmarshal '%s'", string(terms))
		}
		return termsStruct.Buckets, nil
	}
	return nil, errors.Errorf("no field '%s' in facet", key)
}

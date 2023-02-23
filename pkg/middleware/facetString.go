package middleware

import (
	"emperror.dev/errors"
	"encoding/json"
	"fmt"
	"github.com/je4/elasticdsl/v2/pkg/dsl"
	"github.com/je4/elasticdsl/v2/pkg/elastic"
)

type StringFacet struct {
	Name   string
	Values []string
}

type StringFacetResult struct {
	DocCount int    `json:"doc_count"`
	Key      string `json:"key"`
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

func (sf StringFacet) FromJSON(dataBytes []byte) ([]*StringFacetResult, error) {
	var data = map[string]elastic.JSONDummy{}
	if err := json.Unmarshal(dataBytes, &data); err != nil {
		return nil, errors.Wrapf(err, "cannot unmarshal '%s'", string(dataBytes))
	}
	var sfr = &StringFacetsResult{}
	if docCountBytes, ok := data["doc_count"]; ok {
		if err := json.Unmarshal(docCountBytes, &sfr.DocCount); err != nil {
			return nil, errors.Wrapf(err, "cannot unmarshal '%s'", string(docCountBytes))
		}
	}
	key := fmt.Sprintf("facet-strings-%s-terms", sf.GetName())
	if terms, ok := data[key]; ok {
		termsStruct := &struct {
			DocCountErrorUpperBound int                  `json:"doc_count_error_upper_bound,omitempty"`
			SumOtherDocCount        int                  `json:"sum_other_doc_count,omitempty"`
			Buckets                 []*StringFacetResult `json:"buckets"`
		}{}
		if err := json.Unmarshal(terms, termsStruct); err != nil {
			return nil, errors.Wrapf(err, "cannot unmarshal '%s'", string(terms))
		}
		return termsStruct.Buckets, nil
	}
	return nil, errors.Errorf("no field '%s' in facet", key)
}
